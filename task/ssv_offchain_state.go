package task

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

const (
	// todo: automatically detect event fetch limit from target rpc
	fetchEventBlockLimit      = uint64(4900) * 7
	fetchEth1WaitBlockNumbers = uint64(2)

	eventNameValidatorAdded             = "ValidatorAdded"             // 'ValidatorAdded',
	eventNameValidatorRemoved           = "ValidatorRemoved"           // 'ValidatorRemoved',
	eventNameClusterDeposited           = "ClusterDeposited"           // 'ClusterDeposited',
	eventNameClusterWithdrawn           = "ClusterWithdrawn"           // 'ClusterWithdrawn',
	eventNameClusterLiquidated          = "ClusterLiquidated"          // 'ClusterLiquidated',
	eventNameClusterReactivated         = "ClusterReactivated"         // 'ClusterReactivated',
	eventNameFeeRecipientAddressUpdated = "FeeRecipientAddressUpdated" // 'FeeRecipientAddressUpdated',
)

func (task *Task) updateSsvOffchainState() (retErr error) {
	latestBlockNumber, err := task.connectionOfSuperNodeAccount.Eth1LatestBlock()
	if err != nil {
		return err
	}

	if latestBlockNumber > fetchEth1WaitBlockNumbers {
		latestBlockNumber -= fetchEth1WaitBlockNumbers
	}

	logrus.Debugf("latestBlockNumber: %d, dealedBlockNumber: %d", latestBlockNumber, task.dealedEth1Block)

	if latestBlockNumber <= uint64(task.dealedEth1Block) {
		return nil
	}

	start := uint64(task.dealedEth1Block + 1)
	end := latestBlockNumber

	for i := start; i <= end; i += fetchEventBlockLimit {
		subStart := i
		subEnd := i + fetchEventBlockLimit - 1
		if end < i+fetchEventBlockLimit {
			subEnd = end
		}

		topics, err := utils.EventTopics(task.ssvNetworkAbi, eventNameValidatorAdded, eventNameValidatorRemoved, eventNameClusterDeposited,
			eventNameClusterWithdrawn, eventNameClusterLiquidated, eventNameClusterReactivated, eventNameFeeRecipientAddressUpdated)
		if err != nil {
			return err
		}

		ssvAddress := task.ssvKeyPair.CommonAddress()
		ssvAddressTopic := common.Hash{}
		copy(ssvAddressTopic[common.HashLength-common.AddressLength:], ssvAddress[:])

		logs, err := task.eth1Client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(subStart),
			ToBlock:   new(big.Int).SetUint64(subEnd),
			Addresses: []common.Address{task.ssvNetworkContractAddress},
			Topics:    [][]common.Hash{topics, {ssvAddressTopic}},
		})
		if err != nil {
			return err
		}

		for _, log := range logs {
			if log.Removed {
				logrus.Warnf("log %s was removed, will skip", log.TxHash.String())
				continue
			}

			event, err := task.ssvNetworkAbi.EventByID(log.Topics[0])
			if err != nil {
				return err
			}

			switch event.RawName {
			case eventNameValidatorAdded:
				event := &ssv_network.SsvNetworkValidatorAdded{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameValidatorAdded, log.Data, log.Topics)
				if err != nil {
					return err
				}

				pubkey := hex.EncodeToString(event.PublicKey)

				logrus.Debugf("find event validatorAddded, block %d, val: %s tx: %s", event.Raw.BlockNumber, pubkey, event.Raw.TxHash.String())

				// set val's cluster key
				val, exist := task.validatorsByPubkey[pubkey]
				if !exist {
					return fmt.Errorf("val %s not exist in offchain state", pubkey)
				}
				cltKey := clusterKey(event.OperatorIds)
				val.clusterKey = cltKey

				// update cluster
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameValidatorAdded, val.keyIndex)
				if err != nil {
					return err
				}

			case eventNameValidatorRemoved:
				event := &ssv_network.SsvNetworkValidatorRemoved{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameValidatorRemoved, log.Data, log.Topics)
				if err != nil {
					return err
				}

				pubkey := hex.EncodeToString(event.PublicKey)

				logrus.Debugf("find event validatorRemoved, block %d, val: %s tx: %s", event.Raw.BlockNumber, pubkey, event.Raw.TxHash.String())

				val, exist := task.validatorsByPubkey[pubkey]
				if !exist {
					return fmt.Errorf("val %s not exist in offchain state", pubkey)
				}

				// update removed block
				if val.removedFromSsvOnBlock < event.Raw.BlockNumber {
					val.removedFromSsvOnBlock = event.Raw.BlockNumber
				}

				// update cluster
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameValidatorRemoved, val.keyIndex)
				if err != nil {
					return err
				}

			case eventNameClusterDeposited:
				event := &ssv_network.SsvNetworkClusterDeposited{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameClusterDeposited, log.Data, log.Topics)
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterDeposited, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterDeposited, 0)
				if err != nil {
					return err
				}
			case eventNameClusterWithdrawn:
				event := &ssv_network.SsvNetworkClusterWithdrawn{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameClusterWithdrawn, log.Data, log.Topics)
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterWithdrawn, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterWithdrawn, 0)
				if err != nil {
					return err
				}
			case eventNameClusterLiquidated:
				event := &ssv_network.SsvNetworkClusterLiquidated{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameClusterLiquidated, log.Data, log.Topics)
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterLiquidated, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterLiquidated, 0)
				if err != nil {
					return err
				}
			case eventNameClusterReactivated:
				event := &ssv_network.SsvNetworkClusterReactivated{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameClusterReactivated, log.Data, log.Topics)
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterReactivated, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterReactivated, 0)
				if err != nil {
					return err
				}
			case eventNameFeeRecipientAddressUpdated:
				event := &ssv_network.SsvNetworkFeeRecipientAddressUpdated{Raw: log}
				err := utils.UnpackEvent(task.ssvNetworkAbi, event, eventNameFeeRecipientAddressUpdated, log.Data, log.Topics)
				if err != nil {
					return err
				}

				logrus.Debugf("find event FeeRecipientAddressIter, tx: %s, feeRecipient: %s", event.Raw.TxHash.String(),
					event.RecipientAddress.String())
				task.feeRecipientAddressOnSsv = event.RecipientAddress
			default:
				return fmt.Errorf("unknown event, tx: %s", hex.EncodeToString(log.TxHash[:]))
			}
		}

		task.dealedEth1Block = subEnd

		logrus.WithFields(logrus.Fields{
			"start": subStart,
			"end":   subEnd,
		}).Debug("already dealed blocks")
	}

	logrus.WithFields(logrus.Fields{
		"feeRecipientAddressOnSsv": task.feeRecipientAddressOnSsv,
	}).Debug("offchain-state")

	// update cluster balance
	for key, c := range task.clusters {
		balance, err := task.ssvNetworkViewsContract.GetBalance(nil, task.ssvKeyPair.CommonAddress(), c.operatorIds, ssv_network_views.ISSVNetworkCoreCluster(*c.latestCluster))
		if err != nil {
			if strings.Contains(err.Error(), "execution reverted") {
				balance = big.NewInt(0)
			} else {
				return errors.Wrap(err, "ssvNetworkViewsContract.GetBalance failed")
			}
		}
		c.balance = decimal.NewFromBigInt(balance, 0)

		logrus.WithFields(logrus.Fields{
			"clusterKey":                        key,
			"operators":                         c.operatorIds,
			"latestState":                       c.latestCluster,
			"managingValidators":                c.managingValidators,
			"latestUpdateClusterBlockNumber":    c.latestUpdateClusterBlockNumber,
			"latestValidatorAddedBlockNumber":   c.latestValidatorAddedBlockNumber,
			"latestValidatorRemovedBlockNumber": c.latestValidatorRemovedBlockNumber,
			"balance":                           c.balance.String(),
		}).Debug("clusterInfo")

	}
	return nil
}

func (task *Task) updateCluster(operatorIds []uint64, newCluster *ssv_network.ISSVNetworkCoreCluster, blockNumber uint64, eventName string, keyIndex int) error {
	sort.Slice(operatorIds, func(i, j int) bool {
		return operatorIds[i] < operatorIds[j]
	})

	cltKey := clusterKey(operatorIds)
	cluster, exist := task.clusters[cltKey]
	if !exist {

		// for _, opId := range operatorIds {
		// 	if _, exist := task.targetOperators[opId]; !exist {

		// 		willUsePubkey := ""
		// 		if pubkey, exist := task.operatorPubkeys[opId]; exist {
		// 			willUsePubkey = pubkey
		// 		}

		// 		// fetch active status from api
		// 		operatorFromApi, err := utils.MustGetOperatorDetail(task.ssvApiNetwork, opId)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		isActive := false
		// 		if operatorFromApi.IsActive == 1 {
		// 			isActive = true
		// 		}

		// 		fee := big.NewInt(0)
		// 		validatorCount := uint32(0)
		// 		if isActive {
		// 			_, fee, validatorCount, _, _, _, err = task.ssvNetworkViewsContract.GetOperatorById(nil, opId)
		// 			if err != nil {
		// 				return errors.Wrap(err, "ssvNetworkViewsContract.GetOperatorById failed")
		// 			}
		// 		}

		// 		task.targetOperators[opId] = &keyshare.Operator{
		// 			Id:             opId,
		// 			PublicKey:      willUsePubkey,
		// 			Fee:            decimal.NewFromBigInt(fee, 0),
		// 			Active:         isActive,
		// 			ValidatorCount: uint64(validatorCount),
		// 		}

		// 	}
		// }

		cluster = &Cluster{
			operatorIds:        operatorIds,
			latestCluster:      newCluster,
			managingValidators: make(map[int]struct{}),
		}

		task.clusters[cltKey] = cluster
	}

	// update cluster
	if cluster.latestUpdateClusterBlockNumber < blockNumber {
		cluster.latestUpdateClusterBlockNumber = blockNumber
		cluster.latestCluster = newCluster
	}

	switch eventName {
	case eventNameValidatorAdded:
		if cluster.latestValidatorAddedBlockNumber < blockNumber {
			cluster.latestValidatorAddedBlockNumber = blockNumber
			cluster.managingValidators[keyIndex] = struct{}{}

			// update nonce
			task.latestRegistrationNonce++
		}
	case eventNameValidatorRemoved:
		if cluster.latestValidatorRemovedBlockNumber < blockNumber {
			cluster.latestValidatorRemovedBlockNumber = blockNumber

			delete(cluster.managingValidators, keyIndex)
		}

	}

	return nil
}
