package task

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func (task *Task) updateOperatorStatus() error {

	valAmountLimit, err := task.ssvNetworkViewsContract.GetValidatorsPerOperatorLimit(nil)
	if err != nil {
		return err
	}

	task.validatorsPerOperatorLimit = uint64(valAmountLimit)

	logrus.Debugf("validatorsPerOperatorLimit %d", task.validatorsPerOperatorLimit)

	for _, op := range task.targetOperators {

		// check val amount limit per operator
		_, operatorFee, validatorCount, _, _, isActive, err := task.ssvNetworkViewsContract.GetOperatorById(nil, op.Id)
		if err != nil {
			return err
		}

		if isActive {
			op.Active = true
		} else {
			op.Active = false
		}

		// update fee
		op.Fee = decimal.NewFromBigInt(operatorFee, 0)

		// update val count
		op.ValidatorCount = uint64(validatorCount)

		logrus.WithFields(logrus.Fields{
			"id":             op.Id,
			"active":         op.Active,
			"fee":            op.Fee,
			"validatorcount": op.ValidatorCount,
			"pubkey":         op.PublicKey,
		}).Debug("operatorInfo")
	}

	return nil
}
