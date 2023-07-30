package task_ssv

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) checkAndSetFeeRecipient() error {
	logrus.Debug("checkAndSetFeeRecipient start -----------")
	defer func() {
		logrus.Debug("checkAndSetFeeRecipient end -----------")
	}()

	if task.feeRecipientAddressOnSsv != task.feeRecipientAddressOnStafi {
		logrus.WithFields(logrus.Fields{
			"feeRecipientAddressOnSsv":   task.feeRecipientAddressOnSsv.String(),
			"feeRecipientAddressOnStafi": task.feeRecipientAddressOnStafi.String(),
		}).Warn("feeRecipient")
		// send tx
		err := task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}
		defer task.connectionOfSsvAccount.UnlockTxOpts()

		tx, err := task.ssvNetworkContract.SetFeeRecipientAddress(task.connectionOfSsvAccount.TxOpts(), task.feeRecipientAddressOnStafi)
		if err != nil {
			return err
		}

		err = utils.WaitTxOkCommon(task.connectionOfSsvAccount.Eth1Client(), tx.Hash())
		if err != nil {
			return err
		}
		task.feeRecipientAddressOnSsv = task.feeRecipientAddressOnStafi
	}
	return nil
}
