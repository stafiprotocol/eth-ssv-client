package task

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
)

// offboard validator from cluster if:
// 0 exit on beacon
// OR
// 1 operator is not active/fee not match
func (task *Task) checkAndOffboardOnSSV() error {
	if !task.offchainStateIsLatest() {
		return nil
	}

	for i := 0; i < task.nextKeyIndex; i++ {
		val, exist := task.validatorsByKeyIndex[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}

		cluster, exist := task.clusters[val.clusterKey]
		if !exist {
			return nil
		}

		shouldOffboard := false

		// validator is invalid from ssv api
		if val.statusOnSsv == valStatusRegistedOnSsvInvalid {
			shouldOffboard = true
		}

		// validator life end
		if val.statusOnStafi == valStatusStaked &&
			(val.statusOnSsv == valStatusRegistedOnSsvValid || val.statusOnSsv == valStatusRegistedOnSsvInvalid) &&
			val.statusOnBeacon == valStatusExitedOnBeacon {
			shouldOffboard = true
		}

		// life no end but some operators is not active
		if val.statusOnStafi == valStatusStaked &&
			(val.statusOnSsv == valStatusRegistedOnSsvValid || val.statusOnSsv == valStatusRegistedOnSsvInvalid) &&
			val.statusOnBeacon == valStatusActiveOnBeacon {

			inActiveNumber := 0
			for _, opId := range cluster.operatorIds {
				if !task.targetOperators[opId].Active {
					inActiveNumber++
					if inActiveNumber >= opInActiveThreshold {
						logrus.Infof("operator: %d is not active will offboard validator: %d", opId, val.validatorIndex)
						shouldOffboard = true
						break
					}
				}
			}
		}

		// life no end but some operator's fee changed
		if val.statusOnStafi == valStatusStaked &&
			(val.statusOnSsv == valStatusRegistedOnSsvValid || val.statusOnSsv == valStatusRegistedOnSsvInvalid) &&
			val.statusOnBeacon == valStatusActiveOnBeacon {

			preSelectedOperators, err := task.preSelectOperators()
			if err != nil {
				return err
			}
			midFee := preSelectedOperators[len(preSelectedOperators)/2].Fee
			for _, opId := range cluster.operatorIds {
				if task.targetOperators[opId].Fee.GreaterThan(midFee) {
					shouldOffboard = true
					break
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
			return errors.Wrap(err, "ssvNetworkContract.RemoveValidator")
		}
		task.connectionOfSsvAccount.UnlockTxOpts()

		logrus.WithFields(logrus.Fields{
			"txHash":      removeTx.Hash(),
			"operaterIds": cluster.operatorIds,
			"pubkey":      hex.EncodeToString(val.privateKey.PublicKey().Marshal()),
		}).Info("offboard-tx")

		err = task.waitTxOk(removeTx.Hash())
		if err != nil {
			return err
		}

		// offboard one validator per cycle
		break
	}

	return nil
}
