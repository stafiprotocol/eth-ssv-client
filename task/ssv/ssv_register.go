package task_ssv

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/bls"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) checkAndRegisterOnSSV() error {

	for i := 0; i < task.nextKeyIndex; i++ {

		val, exist := task.validators[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}

		logrus.WithFields(logrus.Fields{
			"keyIndex": i,
			"pubkey":   hex.EncodeToString(val.privateKey.PublicKey().Marshal()),
			"status":   val.status,
		}).Debug("register-val")

		if val.status != valStatusStaked {
			continue
		}

		// check status on ssv
		active, err := task.ssvNetworkViewsContract.GetValidator(nil, task.ssvKeyPair.CommonAddress(), val.privateKey.PublicKey().Marshal())
		if err != nil {
			// remove when new SSVViews contract is deployed
			if strings.Contains(err.Error(), "execution reverted") {
				active = false
			} else {
				return errors.Wrap(err, "ssvNetworkViewsContract.GetValidator failed")
			}
		}
		if active {
			return fmt.Errorf("validator %s at index %d is active on ssv", val.privateKey.PublicKey().SerializeToHexStr(), val.keyIndex)
		}

		// encrypt share
		encryptShares, err := keyshare.EncryptShares(val.privateKey.Marshal(), task.operators)
		if err != nil {
			return err
		}

		// build payload
		operatorIds := make([]uint64, 0)
		shares := make([]byte, 0)
		pubkeys := make([]byte, 0)
		ssvAmount := task.clusterInitSsvAmount
		for i, op := range task.operators {
			operatorIds = append(operatorIds, uint64(op.Id))

			shareBts, err := base64.StdEncoding.DecodeString(encryptShares[i].EncryptedKey)
			if err != nil {
				return errors.Wrap(err, "EncryptedKey decode failed")
			}
			shares = append(shares, shareBts...)

			pubkeyBts, err := hexutil.Decode(encryptShares[i].PublicKey)
			if err != nil {
				return errors.Wrap(err, "publickey decode failed")
			}
			pubkeys = append(pubkeys, pubkeyBts...)
		}

		// sign with val private key
		data := fmt.Sprintf("%s:%d", task.connectionOfSsvAccount.TxOpts().From.String(), task.latestRegistrationNonce)
		hash := crypto.Keccak256([]byte(data))
		valPrivateKey, err := bls.PrivateKeyFromBytes(val.privateKey.Marshal())
		if err != nil {
			return err
		}
		sigs := valPrivateKey.Sign(hash).Marshal()

		// build shareData
		shareData := append(sigs, pubkeys...)
		shareData = append(shareData, shares...)

		// check cluster state
		isLiquidated, err := task.ssvNetworkViewsContract.IsLiquidated(nil, task.connectionOfSsvAccount.TxOpts().From, operatorIds, ssv_network_views.ISSVNetworkCoreCluster(*task.latestCluster))
		if err != nil {
			return errors.Wrap(err, "ssvNetworkViewsContract.IsLiquidated failed")
		}
		if isLiquidated {
			logrus.WithFields(logrus.Fields{
				"operators": operatorIds,
			}).Warn("cluster is liquidated")
			return nil
		}

		// check ssv allowance
		allowance, err := task.ssvTokenContract.Allowance(nil, task.connectionOfSsvAccount.TxOpts().From, task.ssvNetworkViewsContractAddress)
		if err != nil {
			return err
		}

		if allowance.Cmp(task.clusterInitSsvAmount) < 0 {
			err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
			if err != nil {
				return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
			}
			approveTx, err := task.ssvTokenContract.Approve(task.connectionOfSsvAccount.TxOpts(), task.ssvNetworkContractAddress, ssvAmount)
			if err != nil {
				task.connectionOfSsvAccount.UnlockTxOpts()
				return err
			}
			task.connectionOfSsvAccount.UnlockTxOpts()

			logrus.WithFields(logrus.Fields{
				"txHash":        approveTx.Hash(),
				"approveAmount": ssvAmount.String(),
			}).Info("approve-tx")

			err = utils.WaitTxOkCommon(task.connectionOfSuperNodeAccount.Eth1Client(), approveTx.Hash())
			if err != nil {
				return err
			}
		}

		// send register tx
		err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}

		registerTx, err := task.ssvNetworkContract.RegisterValidator(task.connectionOfSsvAccount.TxOpts(), val.privateKey.PublicKey().Marshal(), operatorIds, shareData, ssvAmount, ssv_network.ISSVNetworkCoreCluster(*task.latestCluster))
		if err != nil {
			task.connectionOfSsvAccount.UnlockTxOpts()
			return errors.Wrap(err, "ssvNetworkContract.RegisterValidator failed")
		}
		task.connectionOfSsvAccount.UnlockTxOpts()

		logrus.WithFields(logrus.Fields{
			"txHash":      registerTx.Hash(),
			"nonce":       task.latestRegistrationNonce,
			"operaterIds": operatorIds,
			"pubkey":      hex.EncodeToString(val.privateKey.PublicKey().Marshal()),
			"ssvAmount":   ssvAmount.String(),
		}).Info("register-tx")

		err = utils.WaitTxOkCommon(task.connectionOfSuperNodeAccount.Eth1Client(), registerTx.Hash())
		if err != nil {
			return err
		}
	}

	return nil
}
