package task_ssv

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

const (
	// todo: automatically detect event fetch limit from target rpc
	fetchEventBlockLimit      = uint64(4900)
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
	logrus.Debug("updateOffchainState start -----------")
	defer func() {
		logrus.Debug("updateOffchainState end -----------")
	}()

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

		eventValidatorAdded, exist := task.ssvNetworkAbi.Events[eventNameValidatorAdded]
		if !exist {
			return fmt.Errorf("eventValidatorAdded not exist")
		}
		eventValidatorRemoved, exist := task.ssvNetworkAbi.Events[eventNameValidatorRemoved]
		if !exist {
			return fmt.Errorf("eventValidatorRemoved not exist")
		}
		eventClusterDeposited, exist := task.ssvNetworkAbi.Events[eventNameClusterDeposited]
		if !exist {
			return fmt.Errorf("eventClusterDeposited not exist")
		}
		eventClusterWithdrawn, exist := task.ssvNetworkAbi.Events[eventNameClusterWithdrawn]
		if !exist {
			return fmt.Errorf("eventClusterWithdrawn not exist")
		}
		eventClusterLiquidated, exist := task.ssvNetworkAbi.Events[eventNameClusterLiquidated]
		if !exist {
			return fmt.Errorf("eventClusterLiquidated not exist")
		}
		eventClusterReactivated, exist := task.ssvNetworkAbi.Events[eventNameClusterReactivated]
		if !exist {
			return fmt.Errorf("eventClusterReactivated not exist")
		}
		eventFeeRecipientAddressUpdated, exist := task.ssvNetworkAbi.Events[eventNameFeeRecipientAddressUpdated]
		if !exist {
			return fmt.Errorf("eventFeeRecipientAddressUpdated not exist")
		}

		ssvAddress := task.ssvKeyPair.CommonAddress()
		ssvAddressTopic := common.Hash{}
		copy(ssvAddressTopic[common.HashLength-common.AddressLength:], ssvAddress[:])

		logs, err := task.eth1Client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(subStart),
			ToBlock:   new(big.Int).SetUint64(subEnd),
			Addresses: []common.Address{task.ssvNetworkContractAddress},
			Topics: [][]common.Hash{{eventValidatorAdded.ID, eventValidatorRemoved.ID, eventClusterDeposited.ID,
				eventClusterWithdrawn.ID, eventClusterLiquidated.ID, eventClusterReactivated.ID, eventFeeRecipientAddressUpdated.ID},
				{ssvAddressTopic}},
		})
		if err != nil {
			return err
		}

		for _, log := range logs {
			switch log.Topics[0] {
			case eventValidatorAdded.ID:
				event := &ssv_network.SsvNetworkValidatorAdded{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameValidatorAdded, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameValidatorAdded].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
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

				cluster := task.clusters[cltKey]
				cluster.managingValidators[val.keyIndex] = struct{}{}

			case eventValidatorRemoved.ID:
				event := &ssv_network.SsvNetworkValidatorRemoved{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameValidatorRemoved, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameValidatorRemoved].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
				if err != nil {
					return err
				}

				pubkey := hex.EncodeToString(event.PublicKey)

				logrus.Debugf("find event validatorRemoved, block %d, val: %s tx: %s", event.Raw.BlockNumber, pubkey, event.Raw.TxHash.String())

				cluKey := clusterKey(event.OperatorIds)
				val, exist := task.validatorsByPubkey[pubkey]
				if !exist {
					return fmt.Errorf("val %s not exist in offchain state", pubkey)
				}

				// update cluster
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameValidatorRemoved, val.keyIndex)
				if err != nil {
					return err
				}

				cluster := task.clusters[cluKey]
				delete(cluster.managingValidators, val.keyIndex)
			case eventClusterDeposited.ID:
				event := &ssv_network.SsvNetworkClusterDeposited{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameClusterDeposited, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameClusterDeposited].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterDeposited, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterDeposited, 0)
				if err != nil {
					return err
				}
			case eventClusterWithdrawn.ID:
				event := &ssv_network.SsvNetworkClusterWithdrawn{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameClusterWithdrawn, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameClusterWithdrawn].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterWithdrawn, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterWithdrawn, 0)
				if err != nil {
					return err
				}
			case eventClusterLiquidated.ID:
				event := &ssv_network.SsvNetworkClusterLiquidated{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameClusterLiquidated, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameClusterLiquidated].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterLiquidated, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterLiquidated, 0)
				if err != nil {
					return err
				}
			case eventClusterReactivated.ID:
				event := &ssv_network.SsvNetworkClusterReactivated{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameClusterReactivated, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameClusterReactivated].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
				if err != nil {
					return err
				}

				logrus.Debugf("find event clusterReactivated, tx: %s", event.Raw.TxHash.String())
				err = task.updateCluster(event.OperatorIds, &event.Cluster, event.Raw.BlockNumber, eventNameClusterReactivated, 0)
				if err != nil {
					return err
				}
			case eventFeeRecipientAddressUpdated.ID:
				event := &ssv_network.SsvNetworkFeeRecipientAddressUpdated{Raw: log}
				err := task.ssvNetworkAbi.UnpackIntoInterface(event, eventNameFeeRecipientAddressUpdated, log.Data)
				if err != nil {
					return err
				}
				var indexed abi.Arguments
				for _, arg := range task.ssvNetworkAbi.Events[eventNameFeeRecipientAddressUpdated].Inputs {
					if arg.Indexed {
						indexed = append(indexed, arg)
					}
				}
				err = abi.ParseTopics(event, indexed, log.Topics[1:])
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
			return errors.Wrap(err, "ssvNetworkViewsContract.GetBalance failed")
		}
		c.balance = decimal.NewFromBigInt(balance, 0)

		logrus.WithFields(logrus.Fields{
			"clusterKey":                        key,
			"operators":                         c.operatorIds,
			"latestState":                       c.latestCluster,
			"nonce":                             c.latestRegistrationNonce,
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
		operatorDetails := make([]*keyshare.Operator, len(operatorIds))
		for i, opId := range operatorIds {
			operatorDetail, err := utils.GetOperatorDetail(task.ssvApiNetwork, opId)
			if err != nil {
				return err
			}
			feeDeci, err := decimal.NewFromString(operatorDetail.Fee)
			if err != nil {
				return err
			}
			operatorDetails[i] = &keyshare.Operator{
				Id:        opId,
				PublicKey: operatorDetail.PublicKey,
				Fee:       feeDeci,
			}
		}
		cluster = &Cluster{
			operators:               operatorDetails,
			operatorIds:             operatorIds,
			latestCluster:           newCluster,
			latestRegistrationNonce: 0,
			managingValidators:      make(map[int]struct{}),
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
			cluster.latestRegistrationNonce++

			cluster.managingValidators[keyIndex] = struct{}{}
		}
	case eventNameValidatorRemoved:
		if cluster.latestValidatorRemovedBlockNumber < blockNumber {
			cluster.latestValidatorRemovedBlockNumber = blockNumber

			delete(cluster.managingValidators, keyIndex)
		}

	}

	return nil
}
