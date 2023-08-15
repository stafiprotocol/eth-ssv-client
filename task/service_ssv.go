package task

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func (task *Task) ssvService() {
	logrus.Info("start ssv service")
	retry := 0
	for {
		if retry > utils.RetryLimit {
			utils.ShutdownRequestChannel <- struct{}{}
			return
		}

		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:

			for i, handler := range task.handlers {
				funcName := task.handlersName[i]
				logrus.Debugf("handler %s start.........", funcName)

				err := handler()
				if err != nil {
					logrus.Warnf("handler %s failed: %s, will retry.", funcName, err)
					time.Sleep(utils.RetryInterval * 4)
					retry++
					continue
				}
				logrus.Debugf("handler %s end.........", funcName)
			}

			retry = 0
		}

		time.Sleep(48 * time.Second) // 48 blocks
	}
}
