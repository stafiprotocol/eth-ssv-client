package task_ssv

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/constants"
	"github.com/stafiprotocol/eth-ssv-client/pkg/credential"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) initValNextKeyIndex() error {
	task.nextKeyIndex = 0
	return task.checkAndRepairValNexKeyIndex()
}

func (task *Task) checkAndRepairValNexKeyIndex() error {
	logrus.Debug("checkAndRepairValNexKeyIndex start -----------")
	defer func() {
		logrus.Debug("checkAndRepairValNe end -----------")
	}()

	retry := 0
	for {
		if retry > utils.RetryLimit {
			return fmt.Errorf("findNextKeyIndex reach retry limit")
		}
		credential, err := credential.NewCredential(task.copySeed(), task.nextKeyIndex, nil, constants.Chain{}, task.eth1WithdrawalAdress)
		if err != nil {
			return err
		}
		pubkey := credential.SigningPK().Marshal()
		pubkeyStatus, err := task.superNodeContract.GetSuperNodePubkeyStatus(nil, pubkey)
		if err != nil {
			logrus.Warnf("GetSuperNodePubkeyStatus err: %s", err.Error())
			time.Sleep(utils.RetryInterval)
			retry++
			continue
		}

		if uint8(pubkeyStatus.Uint64()) == utils.ValidatorStatusUnInitial {
			break
		}

		task.validators[task.nextKeyIndex] = &Validator{
			privateKey: credential.SigningSk,
			status:     uint8(pubkeyStatus.Uint64()),
			keyIndex:   task.nextKeyIndex,
		}

		logrus.WithFields(logrus.Fields{
			"keyIndex":              task.nextKeyIndex,
			"pubkey":                hex.EncodeToString(pubkey),
			"statusOnStafiContract": pubkeyStatus.Uint64(),
		}).Debug("validator key info")

		task.nextKeyIndex++
	}

	return nil
}
