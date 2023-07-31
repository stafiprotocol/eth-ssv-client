package task_ssv

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

const (
	// todo: automatically detect event fetch limit from target rpc
	fetchEventBlockLimit      = uint64(4900)
	fetchEth1WaitBlockNumbers = uint64(2)
)

func (task *Task) updateSsvOffchainState() error {
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
	maxBlock := uint64(0)

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
			if clusterDepositedIter.Event.Raw.BlockNumber > maxBlock {
				cluster := task.clusters[clusterKey(clusterDepositedIter.Event.OperatorIds)]
				cluster.latestCluster = &clusterDepositedIter.Event.Cluster

				maxBlock = clusterDepositedIter.Event.Raw.BlockNumber
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
			if clusterWithdrawnIter.Event.Raw.BlockNumber > maxBlock {
				cluster := task.clusters[clusterKey(clusterWithdrawnIter.Event.OperatorIds)]
				cluster.latestCluster = &clusterWithdrawnIter.Event.Cluster

				maxBlock = clusterWithdrawnIter.Event.Raw.BlockNumber
			}
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
			if validatorRemovedIter.Event.Raw.BlockNumber > maxBlock {
				cluster := task.clusters[clusterKey(validatorRemovedIter.Event.OperatorIds)]
				cluster.latestCluster = &validatorRemovedIter.Event.Cluster

				maxBlock = validatorRemovedIter.Event.Raw.BlockNumber
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
			cluster := task.clusters[clusterKey(validatorAdddedIter.Event.OperatorIds)]
			cluster.latestRegistrationNonce++

			if validatorAdddedIter.Event.Raw.BlockNumber > maxBlock {
				cluster.latestCluster = &validatorAdddedIter.Event.Cluster

				maxBlock = validatorAdddedIter.Event.Raw.BlockNumber
			}
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
			if clusterLiquidatedIter.Event.Raw.BlockNumber > maxBlock {
				cluster := task.clusters[clusterKey(clusterLiquidatedIter.Event.OperatorIds)]
				cluster.latestCluster = &clusterLiquidatedIter.Event.Cluster

				maxBlock = clusterLiquidatedIter.Event.Raw.BlockNumber
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
			if clusterReactivatedIter.Event.Raw.BlockNumber > maxBlock {
				cluster := task.clusters[clusterKey(clusterReactivatedIter.Event.OperatorIds)]
				cluster.latestCluster = &clusterReactivatedIter.Event.Cluster

				maxBlock = clusterReactivatedIter.Event.Raw.BlockNumber
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
