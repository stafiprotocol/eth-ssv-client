package task_ssv

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/eth/v1"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/types"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) updateValStatus() error {
	logrus.Debug("updateValStatus start -----------")
	defer func() {
		logrus.Debug("updateValStatus end -----------")
	}()

	for i := 0; i < task.nextKeyIndex; i++ {
		val, exist := task.validatorsByKeyIndex[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}
		if val.status == valStatusRemovedOnSsv {
			continue
		}

		// status on stafi contract
		if val.status < valStatusStaked {
			pubkeyStatus, err := task.mustGetSuperNodePubkeyStatus(val.privateKey.PublicKey().Marshal())
			if err != nil {
				return fmt.Errorf("mustGetSuperNodePubkeyStatus err: %s", err.Error())
			}

			switch pubkeyStatus {
			case utils.ValidatorStatusUnInitial:
				return fmt.Errorf("validator %s at index %d not exist on chain", val.privateKey.PublicKey().SerializeToHexStr(), i)
			case utils.ValidatorStatusDeposited:
				val.status = valStatusDeposited
			case utils.ValidatorStatusWithdrawMatch:
				val.status = valStatusMatch
			case utils.ValidatorStatusWithdrawUnmatch:
				val.status = valStatusUnmatch
			case utils.ValidatorStatusStaked:
				val.status = valStatusStaked
			default:
				return fmt.Errorf("validator %s at index %d unknown status %d", val.privateKey.PublicKey().SerializeToHexStr(), i, pubkeyStatus)
			}
		}

		// status on ssv contract
		if val.status == valStatusStaked {
			active, err := task.ssvNetworkViewsContract.GetValidator(nil, task.ssvKeyPair.CommonAddress(), val.privateKey.PublicKey().Marshal())
			if err != nil {
				// remove when new SSVViews contract is deployed
				if strings.Contains(err.Error(), "execution reverted") {
					active = false
				} else {
					return errors.Wrap(err, "ssvNetworkViewsContract.GetValidator failed")
				}
			}

			if active {
				val.status = valStatusRegistedOnSsv
			}
		}

		// status on beacon
		continue
		if val.status == valStatusRegistedOnSsv || val.status == valStatusActiveOnBeacon {
			beaconHead, err := task.connectionOfSuperNodeAccount.Eth2BeaconHead()
			if err != nil {
				return errors.Wrap(err, "connectionOfSuperNodeAccount.Eth2BeaconHead failed")
			}

			valStatus, err := task.connectionOfSuperNodeAccount.GetValidatorStatus(types.BytesToValidatorPubkey(val.privateKey.PublicKey().Marshal()),
				&beacon.ValidatorStatusOptions{
					Epoch: &beaconHead.Epoch,
				})
			if err != nil {
				return err
			}

			if !valStatus.Exists {
				continue
			}
			if valStatus.Index == 0 {
				return fmt.Errorf("val %s index is zero", hex.EncodeToString(val.privateKey.PublicKey().Marshal()))
			}

			logrus.WithFields(logrus.Fields{
				"validator": hex.EncodeToString(val.privateKey.PublicKey().Marshal()),
				"status":    valStatus,
			}).Debug("valStatus")

			// cache validator by val index
			task.validatorsByValIndexMutex.Lock()
			if _, exist := task.validatorsByValIndex[valStatus.Index]; !exist {
				val.validatorIndex = valStatus.Index
				task.validatorsByValIndex[valStatus.Index] = val
			}
			task.validatorsByValIndexMutex.Unlock()

			switch valStatus.Status {
			case ethpb.ValidatorStatus_PENDING_INITIALIZED, ethpb.ValidatorStatus_PENDING_QUEUED: // pending
			case ethpb.ValidatorStatus_ACTIVE_ONGOING, ethpb.ValidatorStatus_ACTIVE_EXITING, ethpb.ValidatorStatus_ACTIVE_SLASHED: // active

				val.status = valStatusActiveOnBeacon

			case ethpb.ValidatorStatus_EXITED_UNSLASHED, ethpb.ValidatorStatus_EXITED_SLASHED, // exited
				ethpb.ValidatorStatus_WITHDRAWAL_POSSIBLE, // withdrawable
				ethpb.ValidatorStatus_WITHDRAWAL_DONE:     // withdrawdone

				val.status = valStatusExitedOnBeacon

			default:
				return fmt.Errorf("unsupported validator status %d", valStatus.Status)
			}
		}
	}

	return nil
}
