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

	logrus.Debugf("fetchNewCluster operators len %d", len(operators.Operators))
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

		_, _, _, _, isPrivate, isActive, err := task.ssvNetworkViewsContract.GetOperatorById(nil, uint64(op.ID))
		if err != nil {
			return err
		}
		if isPrivate {
			continue
		}
		if !isActive {
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

	if len(selectedOperator) != clusterOpAmount {
		return fmt.Errorf("not select enough operators for cluster")
	}

	sort.Slice(operatorIds, func(i, j int) bool {
		return operatorIds[i] < operatorIds[j]
	})
	sort.Slice(selectedOperator, func(i, j int) bool {
		return selectedOperator[i].Id < selectedOperator[j].Id
	})

	logrus.WithFields(logrus.Fields{
		"ids":       operatorIds,
		"operators": selectedOperator,
	}).Debug("final selected operators", operatorIds)

	for _, op := range selectedOperator {
		logrus.WithFields(logrus.Fields{
			"operator": *op,
		}).Debug("operatorDetail")
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
			Active:          true,
			Balance:         big.NewInt(0),
		},
		managingValidators: make(map[int]struct{}),
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

func (task *Task) mustGetOperatorDetail(network string, id uint64) (*utils.OperatorDetail, error) {
	retry := 0
	var operatorDetail *utils.OperatorDetail
	var err error
	for {
		if retry > utils.RetryLimit {
			return nil, fmt.Errorf("GetOperatorDetail reach retry limit")
		}
		operatorDetail, err = utils.GetOperatorDetail(network, id)
		if err != nil {
			logrus.Warnf("GetOperatorDetail err: %s", err.Error())
			time.Sleep(utils.RetryInterval)
			retry++
			continue
		}
		break
	}

	return operatorDetail, nil
}

func (task *Task) mustGetValidator(network, pubkey string) (*utils.SsvValidator, error) {
	retry := 0
	var val *utils.SsvValidator
	var err error
	for {
		if retry > utils.RetryLimit {
			return nil, fmt.Errorf("GetValidator reach retry limit")
		}
		val, err = utils.GetValidator(network, pubkey)
		if err != nil {
			logrus.Warnf("GetValidator err: %s, will retry.", err.Error())
			time.Sleep(utils.RetryInterval)
			retry++
			continue
		}
		break
	}

	return val, nil
}
