package task

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

var valAmountThreshold = uint64(5)
var clusterOpAmount = 4
var opInActiveThreshold = 1

func clusterKey(operators []uint64) string {
	key := strings.Builder{}
	for _, operator := range operators {
		key.WriteString(fmt.Sprintf("%d/", operator))
	}
	return key.String()
}

func (task *Task) preSelectOperators() ([]*keyshare.Operator, error) {
	preSelectedOperators := make([]*keyshare.Operator, 0)
	for _, operator := range task.targetOperators {
		if len(operator.PublicKey) == 0 {
			continue
		}

		if !operator.Active {
			continue
		}
		if operator.ValidatorCount+valAmountThreshold > task.validatorsPerOperatorLimit {
			continue
		}

		preSelectedOperators = append(preSelectedOperators, operator)
	}

	if len(preSelectedOperators) < clusterOpAmount*2 {
		return nil, fmt.Errorf("preSelectedOperators number %d less than %d", len(preSelectedOperators), clusterOpAmount*2)
	}

	sort.Slice(preSelectedOperators, func(i, j int) bool {
		if preSelectedOperators[i].Fee.Equal(preSelectedOperators[j].Fee) {
			return preSelectedOperators[i].Id < preSelectedOperators[j].Id
		} else {
			return preSelectedOperators[i].Fee.LessThan(preSelectedOperators[j].Fee)
		}
	})

	return preSelectedOperators, nil
}

// fetch new cluster and cache
// operators sort by fee asc
func (task *Task) fetchNewClusterAndSave() error {
	preSelectedOperators, err := task.preSelectOperators()
	if err != nil {
		return err
	}
	midFee := preSelectedOperators[len(preSelectedOperators)/2].Fee

	selectedOperatorIds := make([]uint64, 0)
	for _, operator := range preSelectedOperators {
		if operator.Fee.GreaterThan(midFee) {
			break
		}
		// check enough and reset if already exist
		if len(selectedOperatorIds) == clusterOpAmount {
			cltKey := clusterKey(selectedOperatorIds)
			if _, exist := task.clusters[cltKey]; exist {
				selectedOperatorIds = make([]uint64, 0)
				continue
			}

			break
		}

		selectedOperatorIds = append(selectedOperatorIds, operator.Id)
	}

	if len(selectedOperatorIds) != clusterOpAmount {
		return fmt.Errorf("not select enough operators for cluster, please check available operators")
	}

	sort.Slice(selectedOperatorIds, func(i, j int) bool {
		return selectedOperatorIds[i] < selectedOperatorIds[j]
	})

	logrus.WithFields(logrus.Fields{
		"ids": selectedOperatorIds,
	}).Debug("final selected operators")

	cltKey := clusterKey(selectedOperatorIds)
	task.clusters[cltKey] = &Cluster{
		operatorIds:                       selectedOperatorIds,
		latestUpdateClusterBlockNumber:    0,
		latestValidatorAddedBlockNumber:   0,
		latestValidatorRemovedBlockNumber: 0,
		latestCluster: &ssv_network.ISSVNetworkCoreCluster{
			ValidatorCount:  0,
			NetworkFeeIndex: 0,
			Index:           0,
			Active:          true,
			Balance:         big.NewInt(0),
		},
		managingValidators: make(map[int]struct{}),
	}

	return nil
}

func (task *Task) selectLocalClusterForRegister() ([]*Cluster, error) {
	preSelectOperators, err := task.preSelectOperators()
	if err != nil {
		return nil, err
	}
	midFee := preSelectOperators[len(preSelectOperators)/2].Fee

	clusterSelected := make([]*Cluster, 0)
Clusters:
	for _, c := range task.clusters {
		for _, opId := range c.operatorIds {
			operator, exist := task.targetOperators[opId]
			if !exist {
				return nil, fmt.Errorf("operator %d not exist in target operators", opId)
			}
			// check pubkey
			if len(operator.PublicKey) == 0 {
				continue Clusters
			}

			// check val amount limit per operator
			if operator.ValidatorCount+valAmountThreshold > task.validatorsPerOperatorLimit {
				continue Clusters
			}

			// check active
			if !operator.Active {
				continue Clusters
			}
			// check fee
			if operator.Fee.GreaterThan(midFee) {
				continue Clusters
			}
		}

		clusterSelected = append(clusterSelected, c)
	}

	return clusterSelected, nil
}

func (task *Task) selectClusterForRegister() (*Cluster, error) {
	clusterSelectedFirst, err := task.selectLocalClusterForRegister()
	if err != nil {
		return nil, err
	}

	if len(clusterSelectedFirst) == 0 {
		err := task.fetchNewClusterAndSave()
		if err != nil {
			return nil, err
		}
	}

	clusterSelectedFinal, err := task.selectLocalClusterForRegister()
	if err != nil {
		return nil, err
	}

	if len(clusterSelectedFinal) == 0 {
		return nil, fmt.Errorf("select cluster faield, please check target operators")
	}

	sort.Slice(clusterSelectedFinal, func(i, j int) bool {
		return len(clusterSelectedFinal[i].managingValidators) < len(clusterSelectedFinal[j].managingValidators)
	})

	return clusterSelectedFinal[0], nil
}

func (task *Task) mustGetSuperNodePubkeyStatus(pubkey []byte) (uint8, error) {
	retry := 0
	var pubkeyStatus *big.Int
	var err error
	for {
		if retry > utils.RetryLimit {
			return 0, fmt.Errorf("updateValStatus reach retry limit")
		}
		pubkeyStatus, err = task.superNodeContract.GetSuperNodePubkeyStatus(nil, pubkey)
		if err != nil {
			logrus.Warnf("GetSuperNodePubkeyStatus err: %s", err.Error())
			time.Sleep(utils.RetryInterval)
			retry++
			continue
		}
		break
	}

	return uint8(pubkeyStatus.Uint64()), nil
}

func (task *Task) copySeed() []byte {
	copyBts := make([]byte, len(task.seed))
	copy(copyBts, task.seed)
	return copyBts
}

func (task *Task) calClusterNeedDepositAmount(cluster *Cluster) (min, max *big.Int, err error) {
	networkFee, err := task.ssvNetworkViewsContract.GetNetworkFee(nil)
	if err != nil {
		return nil, nil, err
	}
	networkFeeDeci := decimal.NewFromBigInt(networkFee, 0)

	ltp, err := task.ssvNetworkViewsContract.GetLiquidationThresholdPeriod(nil)
	if err != nil {
		return nil, nil, err
	}
	ltpDeci := decimal.NewFromInt(int64(ltp))

	balance, err := task.ssvNetworkViewsContract.GetBalance(nil, task.ssvKeyPair.CommonAddress(), cluster.operatorIds,
		ssv_network_views.ISSVNetworkCoreCluster(*cluster.latestCluster))
	if err != nil {
		if strings.Contains(err.Error(), "execution reverted") {
			balance = big.NewInt(0)
		} else {
			return nil, nil, err
		}
	}
	balanceDeci := decimal.NewFromBigInt(balance, 0)

	totalOpFee := decimal.Zero
	for _, opId := range cluster.operatorIds {
		operator, exist := task.targetOperators[opId]
		if !exist {
			return nil, nil, fmt.Errorf("operator %d not exist in target operators", opId)
		}

		totalOpFee = totalOpFee.Add(operator.Fee)
	}
	totalOpFee = totalOpFee.Add(networkFeeDeci)

	valAmount := decimal.NewFromInt(int64(len(cluster.managingValidators) + 1))

	maxExpected := valAmount.Mul(totalOpFee).Mul(blocksOfOneYear.Add(ltpDeci))
	minExpected := valAmount.Mul(totalOpFee).Mul(ltpDeci.Mul(decimal.NewFromInt(2)))

	if maxExpected.LessThan(minExpected) {
		maxExpected = minExpected
	}

	switch {
	case balanceDeci.GreaterThanOrEqual(maxExpected):
		return big.NewInt(0), big.NewInt(0), nil
	case balanceDeci.GreaterThanOrEqual(minExpected) && balanceDeci.LessThan(maxExpected):
		return big.NewInt(0), maxExpected.Sub(balanceDeci).BigInt(), nil
	case balanceDeci.LessThan(minExpected):
		return maxExpected.Sub(balanceDeci).BigInt(), maxExpected.Sub(balanceDeci).BigInt(), nil
	default:
		return nil, nil, fmt.Errorf("unreached balance")
	}

}

// func unpackOperatorPublicKey(fieldBytes []byte) ([]byte, error) {
// 	abi, err := operator_pubkey.OperatorPubkeyMetaData.GetAbi()
// 	if err != nil {
// 		return nil, err
// 	}
// 	outField, err := abi.Unpack("method", fieldBytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("unpack: %w", err)
// 	}

// 	unpacked, ok := outField[0].([]byte)
// 	if !ok {
// 		return nil, fmt.Errorf("cast OperatorPublicKey to []byte: %w", err)
// 	}

// 	return unpacked, nil
// }
