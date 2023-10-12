package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
)

func (task *Task) checkAndWithdrawOnSSV() error {
	if !task.offchainStateIsLatest() {
		return nil
	}

	for _, cluster := range task.clusters {
		// skip clusters with validators
		if len(cluster.managingValidators) != 0 {
			continue
		}

		if cluster.balance.IsZero() {
			continue
		}

		shouldWithdraw := false
		for _, opId := range cluster.operatorIds {
			operator, exist := task.targetOperators[opId]
			if !exist {
				return fmt.Errorf("operator %d not exist in target operators", opId)
			}
			if !operator.Active {
				logrus.Infof("operator: %d is not active will withdraw cluster: %v", opId, cluster.operatorIds)
				shouldWithdraw = true
				break
			}
		}
		if !shouldWithdraw {
			continue
		}

		// send tx
		err := task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}
		defer task.connectionOfSsvAccount.UnlockTxOpts()

		withdrawTx, err := task.ssvNetworkContract.Withdraw(task.connectionOfSsvAccount.TxOpts(),
			cluster.operatorIds, cluster.balance.BigInt(), ssv_network.ISSVNetworkCoreCluster(*cluster.latestCluster))
		if err != nil {
			return errors.Wrap(err, "ssvNetworkContract.RegisterValidator failed")
		}

		logrus.WithFields(logrus.Fields{
			"txHash":     withdrawTx.Hash(),
			"amount":     cluster.balance.String(),
			"clusterKey": clusterKey(cluster.operatorIds),
		}).Info("withdraw-tx")

		err = task.waitTxOk(withdrawTx.Hash())
		if err != nil {
			return err
		}

		break
	}
	return nil
}
