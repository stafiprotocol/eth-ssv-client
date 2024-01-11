package task

import (
	"fmt"
	"time"

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

	rspOperatorsOnChain, err := utils.BatchGetOperatorsOnChain(task.multicaler, task.ssvNetworkViewsContractAddress, opIds)
	if err != nil {
		return err
	}

	for _, op := range task.targetOperators {
		rspOperatorOnChain, exist := rspOperatorsOnChain[op.Id]
		if !exist {
			return fmt.Errorf("operator : %d not fetch", op.Id)
		}

		if op.Id == 9 {
			op.Active = false
		} else {

			// get active status from api
			rspOperatorFromApi, err := task.mustGetOperatorDetail(task.ssvApiNetwork, op.Id)
			if err != nil {
				return err
			}

			if rspOperatorFromApi.IsActive == 1 {
				op.Active = true
				op.LastNotActiveTime = 0
			} else {
				now := time.Now().Unix()

				if op.LastNotActiveTime == 0 {
					op.Active = true
					op.LastNotActiveTime = now
				} else {
					// 1h
					if (now - op.LastNotActiveTime) > 1*60*60 {
						op.Active = false
					} else {
						op.Active = true
					}
				}
			}

			// update fee
			op.Fee = decimal.NewFromBigInt(rspOperatorOnChain.Fee, 0)

			// update val count
			op.ValidatorCount = uint64(rspOperatorOnChain.ValidatorCount)

		}
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
