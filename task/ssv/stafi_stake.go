package task_ssv

import (
	"encoding/hex"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/credential"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) checkAndStake() error {
	logrus.Debug("checkAndStake start -----------")
	defer func() {
		logrus.Debug("checkAndStake end -----------")
	}()

	poolBalance, err := task.userDepositContract.GetBalance(nil)
	if err != nil {
		return err
	}
	poolBalanceDeci := decimal.NewFromBigInt(poolBalance, 0)

	logrus.WithFields(logrus.Fields{
		"balance": poolBalanceDeci.String(),
	}).Debug("stake-poolBalance")

	if poolBalanceDeci.LessThan(minAmountNeedStake) {
		return nil
	}

	// select validators need stake
	validatorsNeedStake := make([]*Validator, 0)
	for i := 0; i < len(task.validatorsByKeyIndex); i++ {
		val := task.validatorsByKeyIndex[i]
		if val.statusOnStafi != valStatusMatch {
			continue
		}

		if poolBalanceDeci.LessThan(minAmountNeedStake) {
			break
		}

		validatorsNeedStake = append(validatorsNeedStake, val)
		poolBalanceDeci = poolBalanceDeci.Sub(superNodeStakeAmount)
	}
	lengthOfValidatorsNeedStake := len(validatorsNeedStake)

	logrus.WithFields(logrus.Fields{
		"stakeLen": lengthOfValidatorsNeedStake,
	}).Debug("stake-info")

	if lengthOfValidatorsNeedStake == 0 {
		return nil
	}

	// build payload
	validatorPubkeys := make([][]byte, lengthOfValidatorsNeedStake)
	sigs := make([][]byte, lengthOfValidatorsNeedStake)
	dataRoots := make([][32]byte, lengthOfValidatorsNeedStake)
	for i, val := range validatorsNeedStake {
		credential, err := credential.NewCredential(task.seed, val.keyIndex, superNodeStakeAmount.Div(utils.GweiDeci).BigInt(),
			task.chain, task.eth1WithdrawalAdress)
		if err != nil {
			return err
		}

		validatorPubkeys[i] = credential.SigningSk.PublicKey().Marshal()
		depositData, err := credential.SigningDepositData()
		if err != nil {
			return err
		}
		sigBts, err := hex.DecodeString(depositData.Signature)
		if err != nil {
			return err
		}
		sigs[i] = sigBts
		dataRootBts, err := hex.DecodeString(depositData.DepositDataRoot)
		if err != nil {
			return err
		}
		if len(dataRootBts) != 32 {
			return fmt.Errorf("dataRoot length %d not match", len(dataRootBts))
		}
		dataRoots[i] = [32]byte(dataRootBts)
	}

	// send tx
	err = task.connectionOfSuperNodeAccount.LockAndUpdateTxOpts()
	if err != nil {
		return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
	}
	defer task.connectionOfSuperNodeAccount.UnlockTxOpts()

	stakeTx, err := task.superNodeContract.Stake(task.connectionOfSuperNodeAccount.TxOpts(), validatorPubkeys, sigs, dataRoots)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"txHash":           stakeTx.Hash(),
		"validatorPubkeys": validatorPubkeys,
		"sigs":             sigs,
		"dataRoots":        dataRoots,
	}).Info("stake-tx")

	err = utils.WaitTxOkCommon(task.connectionOfSuperNodeAccount.Eth1Client(), stakeTx.Hash())
	if err != nil {
		return err
	}

	return nil
}
