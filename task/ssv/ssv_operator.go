package task_ssv

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) updateOperatorStatus() error {
	logrus.Debug("updateOperatorStatus start -----------")
	defer func() {
		logrus.Debug("updateOperatorStatus end -----------")
	}()

	valAmountLimit, err := task.ssvNetworkViewsContract.GetValidatorsPerOperatorLimit(nil)
	if err != nil {
		return err
	}

	task.validatorsPerOperatorLimit = uint64(valAmountLimit)

	logrus.Debugf("validatorsPerOperatorLimit %d", task.validatorsPerOperatorLimit)

	for _, c := range task.clusters {

		// check val amount limit per operator
		for _, op := range c.operators {
			operatorDetail, err := utils.GetOperatorDetail(task.ssvApiNetwork, op.Id)
			if err != nil {
				return err
			}

			// check active
			if operatorDetail.IsActive != 1 {
				op.Active = false
			} else {
				op.Active = true
			}

			// update fee
			feeDeci, err := decimal.NewFromString(operatorDetail.Fee)
			if err != nil {
				return err
			}
			op.Fee = feeDeci

			// update val count
			op.ValidatorCount = uint64(operatorDetail.ValidatorsCount)

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
