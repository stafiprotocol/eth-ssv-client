package task_ssv

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
	ssv_clusters "github.com/stafiprotocol/eth-ssv-client/bindings/SsvClusters"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

var valAmountThreshold = 5
var clusterOpAmount = 4

// fetch new cluster and cache
func (task *Task) fetchNewCluster() error {
	operators, err := utils.GetOperators(task.ssvApiNetwork)
	if err != nil {
		return err
	}
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
			Id:        uint64(op.ID),
			PublicKey: op.PublicKey,
			Fee:       feeDeci,
		})
		operatorIds = append(operatorIds, uint64(op.ID))
		if len(selectedOperator) == clusterOpAmount {
			cltKey := clusterKey(operatorIds)
			if _, exist := task.clusters[cltKey]; exist {
				selectedOperator = selectedOperator[:len(selectedOperator)-1]
				operatorIds = operatorIds[:len(operatorIds)-1]
				continue
			}

			task.clusters[cltKey] = &Cluster{
				operators:                       selectedOperator,
				operatorIds:                     operatorIds,
				latestUpdateClusterBlockNumber:  0,
				latestUpdateRegisterBlockNumber: 0,
				latestCluster: &ssv_clusters.ISSVNetworkCoreCluster{
					ValidatorCount:  0,
					NetworkFeeIndex: 0,
					Index:           0,
					Active:          false,
					Balance:         big.NewInt(0),
				},
				latestRegistrationNonce: 0,
				managingValidators:      make(map[int]struct{}),
			}

		}
	}

	if len(selectedOperator) != clusterOpAmount {
		return fmt.Errorf("not select enough operators for cluster")
	}
	return nil
}
