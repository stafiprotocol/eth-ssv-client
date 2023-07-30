package task_ssv

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) checkAndReactiveOnSSV() error {
	logrus.Debug("checkAndReactiveOnSSV start -----------")
	defer func() {
		logrus.Debug("checkAndReactiveOnSSV end -----------")
	}()

	operatorIds := make([]uint64, 0)
	for _, op := range task.operators {
		operatorIds = append(operatorIds, uint64(op.Id))
	}
	isLiquidated, err := task.ssvNetworkViewsContract.IsLiquidated(nil, task.connectionOfSsvAccount.TxOpts().From, operatorIds, ssv_network_views.ISSVNetworkCoreCluster(*task.latestCluster))
	if err != nil {
		return errors.Wrap(err, "ssvNetworkViewsContract.IsLiquidated failed")
	}

	if isLiquidated {
		// send tx
		err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}
		defer task.connectionOfSsvAccount.UnlockTxOpts()

		reactiveTx, err := task.ssvNetworkContract.Reactivate(task.connectionOfSsvAccount.TxOpts(), operatorIds, task.clusterInitSsvAmount, ssv_network.ISSVNetworkCoreCluster(*task.latestCluster))
		if err != nil {
			return errors.Wrap(err, "ssvNetworkContract.RegisterValidator failed")
		}

		logrus.WithFields(logrus.Fields{
			"txHash":      reactiveTx.Hash(),
			"operaterIds": operatorIds,
		}).Info("reactive-tx")

		err = utils.WaitTxOkCommon(task.connectionOfSuperNodeAccount.Eth1Client(), reactiveTx.Hash())
		if err != nil {
			return err
		}
	}

	return nil
}
