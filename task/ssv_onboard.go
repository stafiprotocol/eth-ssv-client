package task

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
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

// onboard validator if:
// 0 staked on stafi and uninitial not onboard
// OR
// 1 staked on stafi and active and not onboard
func (task *Task) checkAndOnboardOnSSV() error {
	logrus.Debug("checkAndOnboardOnSSV start -----------")
	defer func() {
		logrus.Debug("checkAndOnboardOnSSV end -----------")
	}()

	for i := 0; i < task.nextKeyIndex; i++ {

		val, exist := task.validatorsByKeyIndex[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}

		if !((val.statusOnStafi == valStatusStaked &&
			val.statusOnBeacon == valStatusUnInitiated &&
			val.statusOnSsv != valStatusRegistedOnSsv) ||
			(val.statusOnStafi == valStatusStaked &&
				val.statusOnBeacon == valStatusActiveOnBeacon &&
				val.statusOnSsv != valStatusRegistedOnSsv)) {
			continue
		}

		// check status on ssv
		onboard, err := task.ssvNetworkViewsContract.GetValidator(nil, task.ssvKeyPair.CommonAddress(), val.privateKey.PublicKey().Marshal())
		if err != nil {
			// remove when new SSVViews contract is deployed
			if strings.Contains(err.Error(), "execution reverted") {
				onboard = false
			} else {
				return errors.Wrap(err, "ssvNetworkViewsContract.GetValidator failed")
			}
		}
		if onboard {
			return fmt.Errorf("validator %s at index %d is onboard on ssv", val.privateKey.PublicKey().SerializeToHexStr(), val.keyIndex)
		}

		cluster, err := task.selectClusterForRegister()
		if err != nil {
			return errors.Wrap(err, "selectClusterForRegister failed")
		}

		// encrypt share
		encryptShares, err := keyshare.EncryptShares(val.privateKey.Marshal(), cluster.operators)
		if err != nil {
			return err
		}

		// build payload
		shares := make([]byte, 0)
		pubkeys := make([]byte, 0)

		_, needDepositAmount, err := task.calClusterNeedDepositAmount(cluster)
		if err != nil {
			return err
		}
		for i := range cluster.operators {
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
		data := fmt.Sprintf("%s:%d", task.ssvKeyPair.Address(), cluster.latestRegistrationNonce)
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
		isLiquidated, err := task.ssvNetworkViewsContract.IsLiquidated(nil, task.ssvKeyPair.CommonAddress(),
			cluster.operatorIds, ssv_network_views.ISSVNetworkCoreCluster(*cluster.latestCluster))
		if err != nil {
			return errors.Wrap(err, "ssvNetworkViewsContract.IsLiquidated failed")
		}
		if isLiquidated {
			logrus.WithFields(logrus.Fields{
				"operators": cluster.operatorIds,
			}).Warn("cluster is liquidated")
			return nil
		}

		// check ssv allowance
		allowance, err := task.ssvTokenContract.Allowance(nil, task.ssvKeyPair.CommonAddress(), task.ssvNetworkViewsContractAddress)
		if err != nil {
			return err
		}

		if allowance.Cmp(needDepositAmount) < 0 {
			err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
			if err != nil {
				return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
			}

			approveAmount := new(big.Int).Mul(needDepositAmount, big.NewInt(10))
			approveTx, err := task.ssvTokenContract.Approve(task.connectionOfSsvAccount.TxOpts(), task.ssvNetworkContractAddress, approveAmount)
			if err != nil {
				task.connectionOfSsvAccount.UnlockTxOpts()
				return err
			}
			task.connectionOfSsvAccount.UnlockTxOpts()

			logrus.WithFields(logrus.Fields{
				"txHash":        approveTx.Hash(),
				"approveAmount": approveAmount.String(),
			}).Info("approve-tx")

			err = utils.WaitTxOkCommon(task.eth1Client, approveTx.Hash())
			if err != nil {
				return err
			}
		}

		// send register tx
		err = task.connectionOfSsvAccount.LockAndUpdateTxOpts()
		if err != nil {
			return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
		}

		registerTx, err := task.ssvNetworkContract.RegisterValidator(task.connectionOfSsvAccount.TxOpts(),
			val.privateKey.PublicKey().Marshal(), cluster.operatorIds, shareData, needDepositAmount, ssv_network.ISSVNetworkCoreCluster(*cluster.latestCluster))
		if err != nil {
			task.connectionOfSsvAccount.UnlockTxOpts()
			return errors.Wrap(err, "ssvNetworkContract.RegisterValidator failed")
		}
		task.connectionOfSsvAccount.UnlockTxOpts()

		logrus.WithFields(logrus.Fields{
			"txHash":           registerTx.Hash(),
			"nonce":            cluster.latestRegistrationNonce,
			"operaterIds":      cluster.operatorIds,
			"pubkey":           hex.EncodeToString(val.privateKey.PublicKey().Marshal()),
			"ssvDepositAmount": needDepositAmount.String(),
		}).Info("onboard-tx")

		err = utils.WaitTxOkCommon(task.eth1Client, registerTx.Hash())
		if err != nil {
			return err
		}

		//todo check validator status on ssv api
	}

	return nil
}
