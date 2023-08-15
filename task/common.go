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

var valAmountThreshold = 5
var clusterOpAmount = 4

func clusterKey(operators []uint64) string {
	key := strings.Builder{}
	for _, operator := range operators {
		key.WriteString(fmt.Sprintf("%d/", operator))
	}
	return key.String()
}

// fetch new cluster and cache
func (task *Task) fetchNewClusterAndSave() error {
	selectedOperator := make([]*keyshare.Operator, 0)
	operatorIds := make([]uint64, 0)
	valAmountLimit, err := task.ssvNetworkViewsContract.GetValidatorsPerOperatorLimit(nil)
	if err != nil {
		return err
	}

	hasTargetOperators := false
	targetOperatorsLen := 0
	// select from target operator
	if len(task.targetOperatorIds) > 0 {
		for _, opId := range task.targetOperatorIds {
			// check enough and reset if already exist
			if len(selectedOperator) == clusterOpAmount {
				cltKey := clusterKey(operatorIds)
				if _, exist := task.clusters[cltKey]; exist {
					selectedOperator = make([]*keyshare.Operator, 0)
					operatorIds = make([]uint64, 0)
					continue
				}

				break
			}

			opDetail, err := task.mustGetOperatorDetail(task.ssvApiNetwork, opId)
			if err != nil {
				return err
			}

			if opDetail.IsActive != 1 {
				continue
			}
			if opDetail.ValidatorsCount > int(valAmountLimit)-valAmountThreshold {
				continue
			}

			_, _, _, _, isPrivate, isActive, err := task.ssvNetworkViewsContract.GetOperatorById(nil, opId)
			if err != nil {
				return err
			}
			if isPrivate {
				continue
			}
			if !isActive {
				continue
			}

			feeDeci, err := decimal.NewFromString(opDetail.Fee)
			if err != nil {
				return err
			}
			selectedOperator = append(selectedOperator, &keyshare.Operator{
				Id:             uint64(opDetail.ID),
				PublicKey:      opDetail.PublicKey,
				Fee:            feeDeci,
				Active:         true,
				ValidatorCount: uint64(opDetail.ValidatorsCount),
			})
			operatorIds = append(operatorIds, uint64(opDetail.ID))
		}
		logrus.Debugf("fetchNewClusterFromTarget operators len %d", len(selectedOperator))

		targetOperatorsLen = len(selectedOperator)
		if targetOperatorsLen > 0 {
			hasTargetOperators = true
		} else {
			return fmt.Errorf("target operators %v unavailable", task.targetOperatorIds)
		}
	}

	// select from api
	operators, err := utils.GetOperators(task.ssvApiNetwork)
	if err != nil {
		return err
	}
	logrus.Debugf("fetchNewClusterFromApi operators len %d", len(operators.Operators))
	for _, op := range operators.Operators {
		// check enough and clear if already exist
		if len(selectedOperator) == clusterOpAmount {
			cltKey := clusterKey(operatorIds)
			if _, exist := task.clusters[cltKey]; exist {
				selectedOperator = selectedOperator[:targetOperatorsLen]
				operatorIds = operatorIds[:targetOperatorsLen]
				continue
			}

			break
		}

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
		hasTargetOperators: hasTargetOperators,
	}

	return nil
}

func (task *Task) selectClusterForRegister() (*Cluster, error) {
	clusterSelected := make([]*Cluster, 0)
	if len(task.targetOperatorIds) > 0 {
		hasUsableTargetOperators := false
	Usable:
		for _, c := range task.clusters {
			if c.hasTargetOperators {
				// check val amount limit per operator
				for _, op := range c.operators {
					if int(op.ValidatorCount) > int(task.validatorsPerOperatorLimit)-valAmountThreshold {
						continue Usable
					}

					if !op.Active {
						continue Usable
					}
				}

				hasUsableTargetOperators = true

				break
			}
		}

		if !hasUsableTargetOperators {
			err := task.fetchNewClusterAndSave()
			if err != nil {
				return nil, err
			}
		}

		// select from target clusters
	TargetOut:
		for _, c := range task.clusters {
			if c.hasTargetOperators {
				// check val amount limit per operator
				for _, op := range c.operators {
					if int(op.ValidatorCount) > int(task.validatorsPerOperatorLimit)-valAmountThreshold {
						continue TargetOut
					}

					if !op.Active {
						continue TargetOut
					}
				}

				clusterSelected = append(clusterSelected, c)
				break
			}
		}

		if len(clusterSelected) == 0 {
			return nil, fmt.Errorf("target operators %v unavailable", task.targetOperatorIds)
		}

		return clusterSelected[0], nil
	}

	// select from api
	if len(task.clusters) == 0 {
		err := task.fetchNewClusterAndSave()
		if err != nil {
			return nil, err
		}
	}
Out:
	for _, c := range task.clusters {
		// check val amount limit per operator
		for _, op := range c.operators {
			if int(op.ValidatorCount) > int(task.validatorsPerOperatorLimit)-valAmountThreshold {
				continue Out
			}

			if !op.Active {
				continue Out
			}
		}

		clusterSelected = append(clusterSelected, c)
	}

	if len(clusterSelected) == 0 {
		err := task.fetchNewClusterAndSave()
		if err != nil {
			return nil, err
		}

	Inner:
		for _, c := range task.clusters {
			// check val amount limit per operator
			for _, op := range c.operators {
				if int(op.ValidatorCount) > int(task.validatorsPerOperatorLimit)-valAmountThreshold {
					continue Inner
				}
			}

			clusterSelected = append(clusterSelected, c)
		}

		if len(clusterSelected) == 0 {
			return nil, fmt.Errorf("selectCluster failed")
		}
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
	for _, op := range cluster.operators {
		totalOpFee = totalOpFee.Add(op.Fee)
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
