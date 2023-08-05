package task_ssv

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

// offboard validator from cluster if:
// 0 exit on beacon
// OR
// 1 operator is not active
func (task *Task) checkAndOffboardOnSSV() error {
	logrus.Debug("checkAndOffboardOnSSV start -----------")
	defer func() {
		logrus.Debug("checkAndOffboardOnSSV end -----------")
	}()

	for i := 0; i < task.nextKeyIndex; i++ {
		val, exist := task.validatorsByKeyIndex[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}

		shouldOffboard := false

		// life end
		if val.statusOnStafi == valStatusStaked &&
			val.statusOnSsv == valStatusRegistedOnSsv &&
			val.statusOnBeacon == valStatusExitedOnBeacon {
			shouldOffboard = true
		}

		cluster := task.clusters[val.clusterKey]
		// life no end but operator not active
		if val.statusOnStafi == valStatusStaked &&
			val.statusOnSsv == valStatusRegistedOnSsv &&
			val.statusOnBeacon == valStatusActiveOnBeacon {

			for _, op := range cluster.operators {
				if !op.Active {
					logrus.Infof("operator: %d is not active will offboard validator: %s", val.validatorIndex)
					shouldOffboard = true
				}
			}
		}
		if !shouldOffboard {
			continue
		}

		// check onboard on ssv
		onboard, err := task.ssvNetworkViewsContract.GetValidator(nil, task.ssvKeyPair.CommonAddress(), val.privateKey.PublicKey().Marshal())
		if err != nil {
			// remove when new SSVViews contract is deployed
			if strings.Contains(err.Error(), "execution reverted") {
				onboard = false
			} else {
				return errors.Wrap(err, "ssvNetworkViewsContract.GetValidator failed")
			}
		}
		if !onboard {
			return fmt.Errorf("validator %s at index %d is offboard on ssv", val.privateKey.PublicKey().SerializeToHexStr(), val.keyIndex)
		}

		// send tx
		err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}
		removeTx, err := task.ssvNetworkContract.RemoveValidator(task.connectionOfSsvAccount.TxOpts(),
			val.privateKey.PublicKey().Marshal(), cluster.operatorIds, ssv_network.ISSVNetworkCoreCluster(*cluster.latestCluster))
		if err != nil {
			task.connectionOfSsvAccount.UnlockTxOpts()
			return err
		}
		task.connectionOfSsvAccount.UnlockTxOpts()

		logrus.WithFields(logrus.Fields{
			"txHash":      removeTx.Hash(),
			"operaterIds": cluster.operatorIds,
			"pubkey":      hex.EncodeToString(val.privateKey.PublicKey().Marshal()),
		}).Info("offboard-tx")

		err = utils.WaitTxOkCommon(task.eth1Client, removeTx.Hash())
		if err != nil {
			return err
		}

		// offboard one validator per cycle
		break
	}

	return nil
}
