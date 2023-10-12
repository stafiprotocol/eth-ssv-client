package task

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) updateOperatorStatus() error {

	valAmountLimit, err := task.ssvNetworkViewsContract.GetValidatorsPerOperatorLimit(nil)
	if err != nil {
		return err
	}

	task.validatorsPerOperatorLimit = uint64(valAmountLimit)

	logrus.Debugf("validatorsPerOperatorLimit %d", task.validatorsPerOperatorLimit)

	opIds := make([]uint64, 0)
	for id := range task.targetOperators {
		opIds = append(opIds, id)
	}

	rspOperators, err := utils.BatchGetOperators(task.multicaler, task.ssvNetworkViewsContractAddress, opIds)
	if err != nil {
		return err
	}

	for _, op := range task.targetOperators {
		rspOperator, exist := rspOperators[op.Id]
		if !exist {
			return fmt.Errorf("operator : %d not fetch", op.Id)
		}

		if rspOperator.IsActive {
			op.Active = true
		} else {
			op.Active = false
		}

		// update fee
		op.Fee = decimal.NewFromBigInt(rspOperator.Fee, 0)

		// update val count
		op.ValidatorCount = uint64(rspOperator.ValidatorCount)

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
