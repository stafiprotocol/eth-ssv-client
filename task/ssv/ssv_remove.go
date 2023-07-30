package task_ssv

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) checkAndRemoveOnSSV() error {
	logrus.Debug("checkAndRemoveOnSSV start -----------")
	defer func() {
		logrus.Debug("checkAndRemoveOnSSV end -----------")
	}()

	for i := 0; i < task.nextKeyIndex; i++ {

		val, exist := task.validators[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}

		if val.status != utils.ValidatorStatusExited {
			continue
		}

		// check status on ssv
		active, err := task.ssvNetworkViewsContract.GetValidator(nil, task.ssvKeyPair.CommonAddress(), val.privateKey.PublicKey().Marshal())
		if err != nil {
			// remove when new SSVViews contract is deployed
			if strings.Contains(err.Error(), "execution reverted") {
				active = false
			} else {
				return errors.Wrap(err, "ssvNetworkViewsContract.GetValidator failed")
			}
		}

		if !active {
			return fmt.Errorf("validator %s at index %d is not active on ssv", val.privateKey.PublicKey().SerializeToHexStr(), val.keyIndex)
		}

		operatorIds := make([]uint64, 0)
		for _, op := range task.operators {
			operatorIds = append(operatorIds, uint64(op.Id))
		}

		// send tx
		err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}
		defer task.connectionOfSsvAccount.UnlockTxOpts()

		removeTx, err := task.ssvNetworkContract.RemoveValidator(task.connectionOfSsvAccount.TxOpts(), val.privateKey.PublicKey().Marshal(), operatorIds, ssv_network.ISSVNetworkCoreCluster(*task.latestCluster))
		if err != nil {
			return err
		}

		err = utils.WaitTxOkCommon(task.connectionOfSuperNodeAccount.Eth1Client(), removeTx.Hash())
		if err != nil {
			return err
		}
	}

	return nil
}
