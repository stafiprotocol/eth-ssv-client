package task

import (
	"fmt"
	"math/big"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) updateOperatorStatus() error {
	if !task.offchainStateIsLatest() {
		return nil
	}

	valAmountLimit, err := task.ssvNetworkViewsContract.GetValidatorsPerOperatorLimit(nil)
	if err != nil {
		return err
	}

	task.validatorsPerOperatorLimit = uint64(valAmountLimit)

	logrus.Debugf("validatorsPerOperatorLimit %d", task.validatorsPerOperatorLimit)

	// append using operators to target operators
	for _, cluster := range task.clusters {
		if len(cluster.managingValidators) > 0 {
			for _, opId := range cluster.operatorIds {
				if _, exist := task.targetOperators[opId]; !exist {

					willUsePubkey := ""
					if pubkey, exist := task.operatorPubkeys[opId]; exist {
						willUsePubkey = pubkey
					}

					// fetch active status from api
					operatorFromApi, err := utils.MustGetOperatorDetail(task.ssvApiNetwork, opId)
					if err != nil {
						return err
					}
					isActive := false
					if operatorFromApi.IsActive == 1 {
						isActive = true
					}

					fee := big.NewInt(0)
					validatorCount := uint32(0)
					if isActive {
						_, fee, validatorCount, _, _, _, err = task.ssvNetworkViewsContract.GetOperatorById(nil, opId)
						if err != nil {
							return errors.Wrap(err, "ssvNetworkViewsContract.GetOperatorById failed")
						}
					}

					task.targetOperators[opId] = &keyshare.Operator{
						Id:             opId,
						PublicKey:      willUsePubkey,
						Fee:            decimal.NewFromBigInt(fee, 0),
						Active:         isActive,
						ValidatorCount: uint64(validatorCount),
					}

				}
			}

		}
	}

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

		// get active status from api
		rspOperatorFromApi, err := utils.MustGetOperatorDetail(task.ssvApiNetwork, op.Id)
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

			logrus.WithFields(logrus.Fields{
				"id":                op.Id,
				"active":            op.Active,
				"LastNotActiveTime": op.LastNotActiveTime,
				"validatorcount":    op.ValidatorCount,
				"pubkey":            op.PublicKey,
			}).Info("get operatorInfo not active info")
		}

		// update fee
		op.Fee = decimal.NewFromBigInt(rspOperatorOnChain.Fee, 0)

		// update val count
		op.ValidatorCount = uint64(rspOperatorOnChain.ValidatorCount)

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
