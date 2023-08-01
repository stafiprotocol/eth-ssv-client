package task_ssv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/signing"
	primTypes "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v3/crypto/bls"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/types"
)

func (task *Task) monitorHandler() {
	logrus.Info("start monitor")

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:
			startCycle, err := task.withdrawContract.EjectedStartCycle(nil)
			if err != nil {
				logrus.Warnf("monitor err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}

			currentCycle, err := task.withdrawContract.CurrentWithdrawCycle(nil)
			if err != nil {
				logrus.Warnf("monitor err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}
			logrus.Debugf("startCycle: %d, currentCycle: %d", startCycle.Uint64(), currentCycle.Uint64())

			start := startCycle.Int64()
			end := currentCycle.Int64()
			if start == 0 {
				start = end - 20
			}
			for i := start; i <= end; {
				err := task.checkCycle(i)
				if err != nil {
					logrus.Warnf("monitor check cycle: %d err: %s", i, err)
					time.Sleep(6 * time.Second)
					continue
				}
				i++
			}
		}

		break
	}

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:

			logrus.Debug("checkCycle start -----------")
			currentCycle, err := task.withdrawContract.CurrentWithdrawCycle(nil)
			if err != nil {
				logrus.Warnf("get currentWithdrawCycle err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}

			start := currentCycle.Int64() - 10
			end := currentCycle.Int64()

			for i := start; i <= end; i++ {
				err = task.checkCycle(i)
				if err != nil {
					logrus.Warnf("checkCycle %d err: %s", i, err)
					time.Sleep(6 * time.Second)
					continue
				}
			}
			logrus.Debug("checkCycle end -----------")

		}

		time.Sleep(60 * time.Second)
	}
}

func (task *Task) uptimeHandler() {

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:
			logrus.Debug("postUptime start -----------")
			err := task.postUptime()
			if err != nil {
				logrus.Warnf("postUptime err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}

			logrus.Debug("postUptime end -----------")
		}

		time.Sleep(5 * time.Minute)
	}
}

func (task *Task) checkCycle(cycle int64) error {
	logrus.Debugf("checkCycle %d", cycle)
	ejectedValidators, err := task.withdrawContract.GetEjectedValidatorsAtCycle(nil, big.NewInt(cycle))
	if err != nil {
		return err
	}

	for _, ejectedValidator := range ejectedValidators {
		if validator, exist := task.validatorsByValIndex[ejectedValidator.Uint64()]; exist {
			logrus.Infof("validator %d elected at cycle %d", validator.validatorIndex, cycle)
			// check beacon sync status
			syncStatus, err := task.connectionOfSuperNodeAccount.Eth2Client().GetSyncStatus()
			if err != nil {
				return err
			}
			if syncStatus.Syncing {
				return errors.New("could not perform exit: beacon node is syncing.")
			}
			beaconHead, err := task.connectionOfSuperNodeAccount.Eth2Client().GetBeaconHead()
			if err != nil {
				return err
			}
			// check exited before
			pubkey := types.BytesToValidatorPubkey(validator.privateKey.PublicKey().Marshal())
			status, err := task.connectionOfSuperNodeAccount.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{Epoch: &beaconHead.Epoch})
			if err != nil {
				return err
			}
			// will skip if already sign exit
			if status.ExitEpoch != math.MaxUint64 {
				logrus.Infof("validator %d will exit at epoch %d", validator.validatorIndex, status.ExitEpoch)
				delete(task.validatorsByValIndex, status.Index)
				continue
			}

			currentEpoch := primTypes.Epoch(beaconHead.Epoch)

			// not active
			if uint64(currentEpoch) < status.ActivationEpoch {
				logrus.Warnf("validator %d is not active and can't exit, will skip, active epoch: %d current epoch: %d", validator.validatorIndex, status.ActivationEpoch, currentEpoch)
				continue
			}
			if currentEpoch < primTypes.Epoch(status.ActivationEpoch)+shardCommitteePeriod {
				logrus.Warnf("validator %d is not active long enough and can't exit, will skip, active epoch: %d current epoch: %d", validator.validatorIndex, status.ActivationEpoch, currentEpoch)
				continue
			}

			// will sign and broadcast exit msg
			exit := &ethpb.VoluntaryExit{Epoch: currentEpoch, ValidatorIndex: primTypes.ValidatorIndex(validator.validatorIndex)}

			domain, err := task.connectionOfSuperNodeAccount.Eth2Client().GetDomainData(domainVoluntaryExit[:], uint64(exit.Epoch))
			if err != nil {
				return errors.Wrap(err, "Get domainData failed")
			}
			exitRoot, err := signing.ComputeSigningRoot(exit, domain)
			if err != nil {
				return errors.Wrap(err, "ComputeSigningRoot failed")
			}

			secretKey, err := bls.SecretKeyFromBytes(validator.privateKey.Marshal())
			if err != nil {
				return errors.Wrap(err, "failed to initialize keys caches from account keystore")
			}
			sig := secretKey.Sign(exitRoot[:])

			err = task.connectionOfSuperNodeAccount.Eth2Client().ExitValidator(validator.validatorIndex, uint64(currentEpoch), types.BytesToValidatorSignature(sig.Marshal()))
			if err != nil {
				return err
			}

			logrus.Infof("validator %d broadcast voluntary exit ok", validator.validatorIndex)

		}
	}
	return nil
}

func (task *Task) postUptime() error {
	valIndexList := make([]uint64, 0)
	for key := range task.validatorsByValIndex {
		valIndexList = append(valIndexList, key)
	}
	req := ReqEjectorUptime{
		ValidatorIndexList: valIndexList,
	}

	jsonValue, err := json.Marshal(req)
	if err != nil {
		return err
	}

	rsp, err := http.Post(task.postUptimeUrl, "application/json", bytes.NewReader(jsonValue))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	rspUptime := RspUptime{}

	err = json.Unmarshal(body, &rspUptime)
	if err != nil {
		return err
	}
	if rspUptime.Status != "80000" {
		return fmt.Errorf("post uptime err: %s", rspUptime.Status)
	}
	return nil
}

type RspUptime struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ReqEjectorUptime struct {
	ValidatorIndexList []uint64 `json:"validatorIndexList"` //hex string list
}
