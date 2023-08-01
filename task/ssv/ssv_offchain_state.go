package task_ssv

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	ssv_clusters "github.com/stafiprotocol/eth-ssv-client/bindings/SsvClusters"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

const (
	// todo: automatically detect event fetch limit from target rpc
	fetchEventBlockLimit      = uint64(4900)
	fetchEth1WaitBlockNumbers = uint64(2)
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

		// ----------- cluster related events

		// 'ClusterDeposited',
		// 'ClusterWithdrawn',
		// 'ValidatorRemoved',
		// 'ValidatorAdded',
		// 'ClusterLiquidated',
		// 'ClusterReactivated',
		clusterDepositedIter, err := task.ssvClustersContract.FilterClusterDeposited(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})

		if err != nil {
			return err
		}
		for clusterDepositedIter.Next() {
			logrus.Debugf("find event clusterDeposited, tx: %s", clusterDepositedIter.Event.Raw.TxHash.String())
			err := task.updateCluster(clusterDepositedIter.Event.OperatorIds, &clusterDepositedIter.Event.Cluster, clusterDepositedIter.Event.Raw.BlockNumber, false)
			if err != nil {
				return err
			}

		}

		clusterWithdrawnIter, err := task.ssvClustersContract.FilterClusterWithdrawn(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})

		if err != nil {
			return err
		}
		for clusterWithdrawnIter.Next() {
			logrus.Debugf("find event clusterWithdrawn, tx: %s", clusterWithdrawnIter.Event.Raw.TxHash.String())
			err := task.updateCluster(clusterWithdrawnIter.Event.OperatorIds, &clusterWithdrawnIter.Event.Cluster, clusterWithdrawnIter.Event.Raw.BlockNumber, false)
			if err != nil {
				return err
			}
		}

		validatorAdddedIter, err := task.ssvClustersContract.FilterValidatorAdded(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})

		if err != nil {
			return err
		}
		for validatorAdddedIter.Next() {
			logrus.Debugf("find event validatorAddded, tx: %s", validatorAdddedIter.Event.Raw.TxHash.String())

			// set val's cluster key
			pubkey := hex.EncodeToString(validatorAdddedIter.Event.PublicKey)
			val, exist := task.validatorsByPubkey[pubkey]
			if !exist {
				return fmt.Errorf("val %s not exist in offchain state", pubkey)
			}
			cltKey := clusterKey(validatorAdddedIter.Event.OperatorIds)
			val.clusterKey = cltKey

			// update cluster
			err := task.updateCluster(validatorAdddedIter.Event.OperatorIds, &validatorAdddedIter.Event.Cluster, validatorAdddedIter.Event.Raw.BlockNumber, true)
			if err != nil {
				return err
			}

			cluster := task.clusters[cltKey]
			cluster.managingValidators[val.keyIndex] = struct{}{}
		}

		validatorRemovedIter, err := task.ssvClustersContract.FilterValidatorRemoved(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})

		if err != nil {
			return err
		}
		for validatorRemovedIter.Next() {
			logrus.Debugf("find event validatorRemoved, tx: %s", validatorRemovedIter.Event.Raw.TxHash.String())
			cluKey := clusterKey(validatorRemovedIter.Event.OperatorIds)
			pubkey := hex.EncodeToString(validatorRemovedIter.Event.PublicKey)
			val, exist := task.validatorsByPubkey[pubkey]
			if !exist {
				return fmt.Errorf("val %s not exist in offchain state", pubkey)
			}

			// update cluster
			err := task.updateCluster(validatorRemovedIter.Event.OperatorIds, &validatorRemovedIter.Event.Cluster, validatorRemovedIter.Event.Raw.BlockNumber, false)
			if err != nil {
				return err
			}

			cluster := task.clusters[cluKey]
			delete(cluster.managingValidators, val.keyIndex)
		}

		clusterLiquidatedIter, err := task.ssvClustersContract.FilterClusterLiquidated(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})

		if err != nil {
			return err
		}
		for clusterLiquidatedIter.Next() {
			logrus.Debugf("find event clusterLiquidated, tx: %s", clusterLiquidatedIter.Event.Raw.TxHash.String())
			err := task.updateCluster(clusterLiquidatedIter.Event.OperatorIds, &clusterLiquidatedIter.Event.Cluster, clusterLiquidatedIter.Event.Raw.BlockNumber, false)
			if err != nil {
				return err
			}
		}

		clusterReactivatedIter, err := task.ssvClustersContract.FilterClusterReactivated(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})

		if err != nil {
			return err
		}
		for clusterReactivatedIter.Next() {
			logrus.Debugf("find event clusterReactivated, tx: %s", clusterReactivatedIter.Event.Raw.TxHash.String())
			err := task.updateCluster(clusterReactivatedIter.Event.OperatorIds, &clusterReactivatedIter.Event.Cluster, clusterReactivatedIter.Event.Raw.BlockNumber, false)
			if err != nil {
				return err
			}
		}

		// -------- feeRecipientAddress

		feeRecipientAddressIter, err := task.ssvNetworkContract.FilterFeeRecipientAddressUpdated(&bind.FilterOpts{
			Start:   subStart,
			End:     &subEnd,
			Context: context.Background(),
		}, []common.Address{task.ssvKeyPair.CommonAddress()})
		if err != nil {
			return err
		}

		for feeRecipientAddressIter.Next() {
			logrus.Debugf("find event FeeRecipientAddressIter, tx: %s, feeRecipient: %s", feeRecipientAddressIter.Event.Raw.TxHash.String(),
				feeRecipientAddressIter.Event.RecipientAddress.String())
			task.feeRecipientAddressOnSsv = feeRecipientAddressIter.Event.RecipientAddress
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

	for key, cluster := range task.clusters {
		logrus.WithFields(logrus.Fields{
			"clusterKey":    key,
			"latestNonce":   cluster.latestRegistrationNonce,
			"latestCluster": cluster.latestCluster,
		}).Debug("offchain-state")
	}
	return nil
}

func (task *Task) updateCluster(operatorIds []uint64, newCluster *ssv_clusters.ISSVNetworkCoreCluster, blockNumber uint64, isValidatorAddedEvent bool) error {
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
			operators:                      operatorDetails,
			operatorIds:                    operatorIds,
			latestUpdateClusterBlockNumber: blockNumber,
			latestCluster:                  newCluster,
			latestRegistrationNonce:        0,
			managingValidators:             make(map[int]struct{}),
		}

		task.clusters[cltKey] = cluster
	} else {
		if cluster.latestUpdateClusterBlockNumber < blockNumber {
			cluster.latestUpdateClusterBlockNumber = blockNumber
			cluster.latestCluster = newCluster
		}
	}

	if isValidatorAddedEvent {
		if cluster.latestUpdateRegisterBlockNumber < blockNumber {
			cluster.latestUpdateRegisterBlockNumber = blockNumber
			cluster.latestRegistrationNonce++
		}
	}

	return nil
}
