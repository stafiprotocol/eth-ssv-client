package task

import (
	"fmt"

	"github.com/pkg/errors"
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

	for _, c := range task.clusters {

		// check val amount limit per operator
		for _, op := range c.operators {
			_, operatorFee, validatorCount, _, isPrivate, _, err := task.ssvNetworkViewsContract.GetOperatorById(nil, op.Id)
			if err != nil {
				return err
			}
			if isPrivate {
				return fmt.Errorf("operator %d is private", op.Id)
			}

			operatorDetail, err := task.mustGetOperatorDetail(task.ssvApiNetwork, op.Id)
			if err != nil {
				return errors.Wrap(err, "mustGetOperatorDetail")
			}

			// check active from api
			if operatorDetail.IsActive == 1 {
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

	}

	return nil
}
