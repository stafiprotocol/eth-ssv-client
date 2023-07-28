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
	fetchEth1WaitBlockNumbers = uint64(5)
)

func (task *Task) updateOffchainState() error {
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

	increaseNonce := func() {
		task.latestRegistrationNonce++
	}

	// 'ClusterDeposited',
	// 'ClusterWithdrawn',
	// 'ValidatorRemoved',
	// 'ValidatorAdded',
	// 'ClusterLiquidated',
	// 'ClusterReactivated',
	for i := start; i <= end; i += fetchEventBlockLimit {
		subStart := i
		subEnd := i + fetchEventBlockLimit - 1
		if end < i+fetchEventBlockLimit {
			subEnd = end
		}

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
				task.latestCluster = &clusterDepositedIter.Event.Cluster
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
				task.latestCluster = &clusterWithdrawnIter.Event.Cluster
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
				task.latestCluster = &validatorRemovedIter.Event.Cluster
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
			increaseNonce()
			if validatorAdddedIter.Event.Raw.BlockNumber > maxBlock {
				task.latestCluster = &validatorAdddedIter.Event.Cluster
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
				task.latestCluster = &clusterLiquidatedIter.Event.Cluster
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
				task.latestCluster = &clusterReactivatedIter.Event.Cluster
				maxBlock = clusterReactivatedIter.Event.Raw.BlockNumber
			}
		}

		task.dealedEth1Block = subEnd

		logrus.WithFields(logrus.Fields{
			"start": subStart,
			"end":   subEnd,
		}).Debug("already dealed blocks")
	}

	logrus.WithFields(logrus.Fields{
		"latestNonce":   task.latestRegistrationNonce,
		"latestCluster": task.latestCluster,
	}).Debug("offchain-state")
	return nil
}
