package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (task *Task) checkAndSetFeeRecipient() error {
	if !task.offchainStateIsLatest() {
		return nil
	}

	if task.feeRecipientAddressOnSsv != task.feeRecipientAddressOnStafi {
		// send tx
		err := task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}
		defer task.connectionOfSsvAccount.UnlockTxOpts()

		tx, err := task.ssvNetworkContract.SetFeeRecipientAddress(task.connectionOfSsvAccount.TxOpts(), task.feeRecipientAddressOnStafi)
		if err != nil {
			return errors.Wrap(err, "ssvNetworkContract.SetFeeRecipientAddress")
		}

		logrus.WithFields(logrus.Fields{
			"feeRecipientAddressOnSsv":   task.feeRecipientAddressOnSsv.String(),
			"feeRecipientAddressOnStafi": task.feeRecipientAddressOnStafi.String(),
			"txHash":                     tx.Hash(),
		}).Warn("setFeeRecipientAddress")

		err = task.waitTxOk(tx.Hash())
		if err != nil {
			return err
		}
		task.feeRecipientAddressOnSsv = task.feeRecipientAddressOnStafi
	}
	return nil
}
