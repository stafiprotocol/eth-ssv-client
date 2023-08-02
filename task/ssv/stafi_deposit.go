package task_ssv

import (
	"encoding/hex"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/credential"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) checkAndDeposit() (retErr error) {
	logrus.Debug("checkAndDeposit start -----------")
	defer func() {
		logrus.Debug("checkAndDeposit end -----------")
	}()

	poolBalance, err := task.userDepositContract.GetBalance(nil)
	if err != nil {
		return err
	}
	poolBalanceDeci := decimal.NewFromBigInt(poolBalance, 0)
	logrus.WithFields(logrus.Fields{
		"balance": poolBalanceDeci.String(),
	}).Debug("deposit-poolBalance")

	if poolBalanceDeci.LessThan(minAmountNeedDeposit) {
		return nil
	}

	// wait if some validators wait stake
	for i := 0; i < len(task.validatorsByKeyIndex); i++ {
		val := task.validatorsByKeyIndex[i]
		if val.status == valStatusDeposited || val.status == valStatusMatch {
			return nil
		}
	}

	// validators need deposit
	depositLen := poolBalanceDeci.Div(minAmountNeedDeposit).IntPart()
	validatorPubkeys := make([][]byte, depositLen)
	sigs := make([][]byte, depositLen)
	dataRoots := make([][32]byte, depositLen)
	oldKeyIndex := task.nextKeyIndex

	logrus.WithFields(logrus.Fields{
		"depositLen":   depositLen,
		"nextKeyIndex": task.nextKeyIndex,
	}).Debug("deposit-info")

	defer func() {
		if retErr != nil {
			// got err, should reset validator info
			task.nextKeyIndex = oldKeyIndex
			for i := task.nextKeyIndex; i < task.nextKeyIndex+int(depositLen); i++ {
				pubkey := hex.EncodeToString(task.validatorsByKeyIndex[i].privateKey.PublicKey().Marshal())

				delete(task.validatorsByKeyIndex, i)
				delete(task.validatorsByPubkey, pubkey)
			}
		}
	}()

	for i := 0; i < int(depositLen); i++ {
		credential, err := credential.NewCredential(task.seed, task.nextKeyIndex, superNodeDepositAmount.Div(utils.GweiDeci).BigInt(), task.chain, task.eth1WithdrawalAdress)
		if err != nil {
			return err
		}

		pubkeyBts := credential.SigningSk.PublicKey().Marshal()
		pubkeyStatus, err := task.mustGetSuperNodePubkeyStatus(pubkeyBts)
		if err != nil {
			return fmt.Errorf("mustGetSuperNodePubkeyStatus err: %s", err.Error())
		}
		if pubkeyStatus != utils.ValidatorStatusUnInitial {
			return fmt.Errorf("pubkey %s at index %d already on chain", hex.EncodeToString(pubkeyBts), task.nextKeyIndex)
		}

		validatorPubkeys[i] = pubkeyBts
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

		val := &Validator{
			privateKey: credential.SigningSk,
			status:     valStatusUnInitiated,
			keyIndex:   task.nextKeyIndex,
		}
		task.validatorsByKeyIndex[task.nextKeyIndex] = val
		task.validatorsByPubkey[hex.EncodeToString(pubkeyBts)] = val

		logrus.WithFields(logrus.Fields{
			"keyIndex":            task.nextKeyIndex,
			"pubkey":              hex.EncodeToString(pubkeyBts),
			"dataRoot":            depositData.DepositDataRoot,
			"messageRoot":         depositData.DepositMessageRoot,
			"sig":                 depositData.Signature,
			"withdrawCredentials": depositData.WithdrawalCredentials,
			"amount":              depositData.Amount,
			"forkVersion":         depositData.ForkVersion,
			"networkName":         depositData.NetworkName,
		}).Info("deposit-params")

		// increase nextKeyIndex
		task.nextKeyIndex++

	}

	err = task.connectionOfSuperNodeAccount.LockAndUpdateTxOpts()
	if err != nil {
		return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
	}
	defer task.connectionOfSuperNodeAccount.UnlockTxOpts()

	depositTx, err := task.superNodeContract.Deposit(task.connectionOfSuperNodeAccount.TxOpts(), validatorPubkeys, sigs, dataRoots)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"txHash":           depositTx.Hash(),
		"validatorPubkeys": validatorPubkeys,
		"sigs":             sigs,
		"dataRoots":        dataRoots,
	}).Info("deposit-tx")

	err = utils.WaitTxOkCommon(task.connectionOfSuperNodeAccount.Eth1Client(), depositTx.Hash())
	if err != nil {
		return err
	}

	for _, pubkey := range validatorPubkeys {
		status, err := task.mustGetSuperNodePubkeyStatus(pubkey)
		if err != nil {
			return fmt.Errorf("mustGetSuperNodePubkeyStatus err: %s", err.Error())
		}
		if status == utils.ValidatorStatusUnInitial {
			return fmt.Errorf("validator %s not exist on chain", hex.EncodeToString(pubkey))
		}
	}

	return nil
}
