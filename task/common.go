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
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

var valAmountThreshold = uint64(5)
var clusterOpAmount = 4
var opInActiveThreshold = 2

func clusterKey(operators []uint64) string {
	key := strings.Builder{}
	for _, operator := range operators {
		key.WriteString(fmt.Sprintf("%d/", operator))
	}
	return key.String()
}

// fetch new cluster and cache
func (task *Task) fetchNewClusterAndSave() error {
	operatorIds := make([]uint64, 0)

	// select from target operator
	for _, operator := range task.targetOperators {
		// check enough and reset if already exist
		if len(operatorIds) == clusterOpAmount {
			cltKey := clusterKey(operatorIds)
			if _, exist := task.clusters[cltKey]; exist {
				operatorIds = make([]uint64, 0)
				continue
			}

			break
		}

		if !operator.Active {
			continue
		}
		if operator.ValidatorCount+valAmountThreshold > task.validatorsPerOperatorLimit {
			continue
		}

		operatorIds = append(operatorIds, operator.Id)
	}

	if len(operatorIds) != clusterOpAmount {
		return fmt.Errorf("not select enough operators for cluster, please check available operators")
	}

	sort.Slice(operatorIds, func(i, j int) bool {
		return operatorIds[i] < operatorIds[j]
	})

	logrus.WithFields(logrus.Fields{
		"ids": operatorIds,
	}).Debug("final selected operators", operatorIds)

	cltKey := clusterKey(operatorIds)
	task.clusters[cltKey] = &Cluster{
		operatorIds:                       operatorIds,
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

func (task *Task) selectClusterForRegister() (*Cluster, error) {
	clusterSelected := make([]*Cluster, 0)
Clusters:
	for _, c := range task.clusters {
		// check val amount limit per operator
		for _, opId := range c.operatorIds {
			operator, exist := task.targetOperators[opId]
			if !exist {
				return nil, fmt.Errorf("operator %d not exist in target operators", opId)
			}

			if operator.ValidatorCount+valAmountThreshold > task.validatorsPerOperatorLimit {
				continue Clusters
			}

			if !operator.Active {
				continue Clusters
			}
		}

		clusterSelected = append(clusterSelected, c)
	}

	if len(clusterSelected) == 0 {
		err := task.fetchNewClusterAndSave()
		if err != nil {
			return nil, err
		}
	}

Clusters2:
	for _, c := range task.clusters {
		// check val amount limit per operator
		for _, opId := range c.operatorIds {
			operator, exist := task.targetOperators[opId]
			if !exist {
				return nil, fmt.Errorf("operator %d not exist in target operators", opId)
			}

			if operator.ValidatorCount+valAmountThreshold > task.validatorsPerOperatorLimit {
				continue Clusters2
			}

			if !operator.Active {
				continue Clusters2
			}
		}

		clusterSelected = append(clusterSelected, c)
	}

	if len(clusterSelected) == 0 {
		return nil, fmt.Errorf("select cluster faield, please check target operators")
	}

	sort.Slice(clusterSelected, func(i, j int) bool {
		return len(clusterSelected[i].managingValidators) < len(clusterSelected[j].managingValidators)
	})

	return clusterSelected[0], nil
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
