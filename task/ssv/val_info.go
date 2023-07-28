package task_ssv

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) updateValStatus() error {
	for i := 0; i < task.nextKeyIndex; i++ {

		val, exist := task.validators[i]
		if !exist {
			return fmt.Errorf("validator at index %d not exist", i)
		}
		if val.status == valStatusRemovedOnSsv {
			continue
		}

		// status on stafi contract
		if val.status < valStatusStaked {
			pubkeyStatus, err := task.mustGetSuperNodePubkeyStatus(val.privateKey.PublicKey().Marshal())
			if err != nil {
				return fmt.Errorf("mustGetSuperNodePubkeyStatus err: %s", err.Error())
			}

			switch pubkeyStatus {
			case utils.ValidatorStatusUnInitial:
				return fmt.Errorf("validator %s at index %d not exist on chain", val.privateKey.PublicKey().SerializeToHexStr(), i)
			case utils.ValidatorStatusDeposited:
				val.status = valStatusDeposited
			case utils.ValidatorStatusWithdrawMatch:
				val.status = valStatusMatch
			case utils.ValidatorStatusWithdrawUnmatch:
				val.status = valStatusUnmatch
			case utils.ValidatorStatusStaked:
				val.status = valStatusStaked
			default:
				return fmt.Errorf("validator %s at index %d unknown status %d", val.privateKey.PublicKey().SerializeToHexStr(), i, pubkeyStatus)
			}
		}

		// status on ssv contract
		if val.status == valStatusStaked {
			active, err := task.ssvNetworkViewsContract.GetValidator(nil, task.ssvKeyPair.CommonAddress(), val.privateKey.PublicKey().Marshal())
			if err != nil {
				// remove when new SSVViews contract is deployed
				if strings.Contains(err.Error(), "execution reverted") {
					active = false
				} else {
					return errors.Wrap(err, "ssvNetworkViewsContract.GetValidator failed")
				}
			}
			if active {
				val.status = valStatusRegistedOnSsv
			}
		}

		task.validators[task.nextKeyIndex] = val
	}
	return nil
}
