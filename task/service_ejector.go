package task

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (task *Task) ejectorService() {
	logrus.Info("start ejector service")

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:
			startCycle, err := task.withdrawContract.EjectedStartCycle(nil)
			if err != nil {
				logrus.Warnf("monitor err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}

			currentCycle, err := task.withdrawContract.CurrentWithdrawCycle(nil)
			if err != nil {
				logrus.Warnf("monitor err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}
			logrus.Debugf("startCycle: %d, currentCycle: %d", startCycle.Uint64(), currentCycle.Uint64())

			start := startCycle.Int64()
			end := currentCycle.Int64()
			if start == 0 {
				start = end - 20
			}
			for i := start; i <= end; {
				err := task.checkCycle(i)
				if err != nil {
					logrus.Warnf("monitor check cycle: %d err: %s", i, err)
					time.Sleep(6 * time.Second)
					continue
				}
				i++
			}
		}

		break
	}

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:

			logrus.Debug("checkCycle start -----------")
			currentCycle, err := task.withdrawContract.CurrentWithdrawCycle(nil)
			if err != nil {
				logrus.Warnf("get currentWithdrawCycle err: %s", err)
				time.Sleep(6 * time.Second)
				continue
			}

			start := currentCycle.Int64() - 10
			end := currentCycle.Int64()

			for i := start; i <= end; i++ {
				err = task.checkCycle(i)
				if err != nil {
					logrus.Warnf("checkCycle %d err: %s", i, err)
					time.Sleep(6 * time.Second)
					continue
				}
			}
			logrus.Debug("checkCycle end -----------")

		}

		time.Sleep(600 * time.Second)
	}
}

func (task *Task) uptimeService() {

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:
			logrus.Debug("postUptime start -----------")
			err := task.postUptime()
			if err != nil {
				logrus.Warnf("postUptime err: %s", err)
				time.Sleep(24 * time.Second)
				continue
			}

			logrus.Debug("postUptime end -----------")
		}

		time.Sleep(5 * time.Minute)
	}
}
