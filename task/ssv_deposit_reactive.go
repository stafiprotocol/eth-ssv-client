package task

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
)

func (task *Task) checkAndReactiveOnSSV() error {
	if !task.offchainStateIsLatest() {
		return nil
	}

	for _, cluster := range task.clusters {
		// skip clusters with zero validators
		if len(cluster.managingValidators) == 0 {
			continue
		}

		// check liquidated
		isLiquidated, err := task.ssvNetworkViewsContract.IsLiquidated(nil, task.ssvKeyPair.CommonAddress(),
			cluster.operatorIds, ssv_network_views.ISSVNetworkCoreCluster(*cluster.latestCluster))
		if err != nil {
			if strings.Contains(err.Error(), "execution reverted") {
				isLiquidated = false
			} else {
				return errors.Wrap(err, "ssvNetworkViewsContract.IsLiquidated failed")
			}
		}

		if isLiquidated {
			// send tx
			err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
			if err != nil {
				return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
			}
			defer task.connectionOfSsvAccount.UnlockTxOpts()

			needDepositAmount, _, err := task.calClusterNeedDepositAmount(cluster)
			if err != nil {
				return err
			}

			reactiveTx, err := task.ssvNetworkContract.Reactivate(task.connectionOfSsvAccount.TxOpts(),
				cluster.operatorIds, needDepositAmount, ssv_network.ISSVNetworkCoreCluster(*cluster.latestCluster))
			if err != nil {
				return errors.Wrap(err, "ssvNetworkContract.RegisterValidator failed")
			}

			logrus.WithFields(logrus.Fields{
				"txHash":     reactiveTx.Hash(),
				"clusterKey": clusterKey(cluster.operatorIds),
			}).Info("reactive-tx")

			err = task.waitTxOk(reactiveTx.Hash())
			if err != nil {
				return err
			}
		} else {

			// check balance and deposit
			needDepositAmount, _, err := task.calClusterNeedDepositAmount(cluster)
			if err != nil {
				return err
			}
			if needDepositAmount.Cmp(big.NewInt(0)) > 0 {
				// send tx
				err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
				if err != nil {
					return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
				}
				defer task.connectionOfSsvAccount.UnlockTxOpts()

				needDepositAmount, _, err := task.calClusterNeedDepositAmount(cluster)
				if err != nil {
					return err
				}

				depositTx, err := task.ssvNetworkContract.Deposit(task.connectionOfSsvAccount.TxOpts(),
					task.ssvKeyPair.CommonAddress(), cluster.operatorIds, needDepositAmount, ssv_network.ISSVNetworkCoreCluster(*cluster.latestCluster))
				if err != nil {
					return errors.Wrap(err, "ssvNetworkContract.Deposit failed")
				}

				logrus.WithFields(logrus.Fields{
					"txHash":     depositTx.Hash(),
					"clusterKey": clusterKey(cluster.operatorIds),
				}).Info("ssv-deposit-tx")

				err = task.waitTxOk(depositTx.Hash())
				if err != nil {
					return err
				}
			}

		}

	}
	return nil
}
