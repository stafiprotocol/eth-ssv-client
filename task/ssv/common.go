package task_ssv

import (
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func clusterKey(operators []uint64) string {
	key := strings.Builder{}
	for _, operator := range operators {
		key.WriteString(fmt.Sprintf("%d/", operator))
	}
	return key.String()
}

var valAmountThreshold = 5
var clusterOpAmount = 4

// fetch new cluster and cache
func (task *Task) fetchNewCluster() error {
	operators, err := utils.GetOperators(task.ssvApiNetwork)
	if err != nil {
		return err
	}

	logrus.Debug("fetchNewCluster operators len %d", len(operators.Operators))
	valAmountLimit, err := task.ssvNetworkViewsContract.GetValidatorsPerOperatorLimit(nil)
	if err != nil {
		return err
	}

	selectedOperator := make([]*keyshare.Operator, 0)
	operatorIds := make([]uint64, 0)
	for _, op := range operators.Operators {
		if op.IsActive != 1 {
			continue
		}
		if op.ValidatorsCount > int(valAmountLimit)-valAmountThreshold {
			continue
		}

		feeDeci, err := decimal.NewFromString(op.Fee)
		if err != nil {
			return err
		}
		selectedOperator = append(selectedOperator, &keyshare.Operator{
			Id:             uint64(op.ID),
			PublicKey:      op.PublicKey,
			Fee:            feeDeci,
			Active:         true,
			ValidatorCount: uint64(op.ValidatorsCount),
		})
		operatorIds = append(operatorIds, uint64(op.ID))

		// check enough and reset if already exist
		if len(selectedOperator) == clusterOpAmount {
			cltKey := clusterKey(operatorIds)
			if _, exist := task.clusters[cltKey]; exist {
				selectedOperator = selectedOperator[:len(selectedOperator)-1]
				operatorIds = operatorIds[:len(operatorIds)-1]
				continue
			}
			break
		}
	}

	logrus.Debug("final selected operators", operatorIds)
	if len(selectedOperator) != clusterOpAmount {
		return fmt.Errorf("not select enough operators for cluster")
	}

	cltKey := clusterKey(operatorIds)
	task.clusters[cltKey] = &Cluster{
		operators:                         selectedOperator,
		operatorIds:                       operatorIds,
		latestUpdateClusterBlockNumber:    0,
		latestValidatorAddedBlockNumber:   0,
		latestValidatorRemovedBlockNumber: 0,
		latestCluster: &ssv_network.ISSVNetworkCoreCluster{
			ValidatorCount:  0,
			NetworkFeeIndex: 0,
			Index:           0,
			Active:          false,
			Balance:         big.NewInt(0),
		},
		latestRegistrationNonce: 0,
		managingValidators:      make(map[int]struct{}),
	}

	return nil
}

func (task *Task) selectClusterForRegister() (*Cluster, error) {
	if len(task.clusters) == 0 {
		err := task.fetchNewCluster()
		if err != nil {
			return nil, err
		}
	}

	clusterSlice := make([]*Cluster, 0)

cluster:
	for _, c := range task.clusters {
		// check val amount limit per operator
		for _, op := range c.operators {
			if int(op.ValidatorCount) > int(task.validatorsPerOperatorLimit)-valAmountThreshold {
				continue cluster
			}

			if !op.Active {
				continue cluster
			}
		}

		clusterSlice = append(clusterSlice, c)
	}

	if len(clusterSlice) == 0 {
		err := task.fetchNewCluster()
		if err != nil {
			return nil, err
		}

	clusterSub:
		for _, c := range task.clusters {
			// check val amount limit per operator
			for _, op := range c.operators {
				if int(op.ValidatorCount) > int(task.validatorsPerOperatorLimit)-valAmountThreshold {
					continue clusterSub
				}
			}

			clusterSlice = append(clusterSlice, c)
		}

		if len(clusterSlice) == 0 {
			return nil, fmt.Errorf("selectCluster failed")
		}
	}

	sort.Slice(clusterSlice, func(i, j int) bool {
		return len(clusterSlice[i].managingValidators) < len(clusterSlice[j].managingValidators)
	})

	return clusterSlice[0], nil
}
