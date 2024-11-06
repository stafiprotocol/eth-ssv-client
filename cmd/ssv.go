// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	"github.com/stafiprotocol/chainbridge/utils/keystore"
	"github.com/stafiprotocol/eth-ssv-client/pkg/config"
	"github.com/stafiprotocol/eth-ssv-client/pkg/log"
	"github.com/stafiprotocol/eth-ssv-client/task"
)

func startSsvCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start ssv client",
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath, err := cmd.Flags().GetString(flagConfigPath)
			if err != nil {
				return err
			}
			fmt.Printf("config path: %s\n", configPath)

			operatorsPath, err := cmd.Flags().GetString(flagOperatorsPath)
			if err != nil {
				return err
			}
			fmt.Printf("operators path: %s\n", operatorsPath)

			logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
			if err != nil {
				return err
			}
			logLevel, err := logrus.ParseLevel(logLevelStr)
			if err != nil {
				return err
			}
			logrus.SetLevel(logLevel)

			cfg, err := config.Load(configPath)
			if err != nil {
				return err
			}
			cfgSub, err := config.Load(operatorsPath)
			if err != nil {
				return err
			}
			if len(cfgSub.Operators) == 0 {
				return fmt.Errorf("operators file: %s is empty", operatorsPath)
			}

			cfg.Operators = cfgSub.Operators

			logrus.Infof(
				`config info:
  logFilePath: %s
  logLevel: %s
  eth1Endpoint: %s
  eth2Endpoint: %s
  storageAddress: %s
  ssvNetworkAddress: %s
  ssvNetworkViewsAddress: %s
  targetOperators: %v
  poolReservedBalance: %s,
  ssvReceiveAddress: %s`,
				cfg.LogFilePath, logLevelStr, cfg.Eth1Endpoint, cfg.Eth2Endpoint,
				cfg.Contracts.StorageContractAddress, cfg.Contracts.SsvNetworkAddress,
				cfg.Contracts.SsvNetworkViewsAddress, cfg.TargetOperators, cfg.PoolReservedBalance, cfg.SsvReceiveAddress)

			err = log.InitLogFile(cfg.LogFilePath + "/ssv")
			if err != nil {
				return err
			}
			if len(cfg.TargetOperators) < 4 {
				return fmt.Errorf("target operators must have at least 4")
			}
			//interrupt signal
			// ctx := utils.ShutdownListener()

			// load super node account
			kpI, err := keystore.KeypairFromAddress(cfg.SuperNodeAccount, keystore.EthChain, cfg.KeystorePath, false)
			if err != nil {
				return err
			}
			kp, ok := kpI.(*secp256k1.Keypair)
			if !ok {
				return fmt.Errorf("super node keypair err")
			}

			// load ssv account
			ssvkpI, err := keystore.KeypairFromAddress(cfg.SsvAccount, keystore.EthChain, cfg.KeystorePath, false)
			if err != nil {
				return err
			}
			ssvkp, ok := ssvkpI.(*secp256k1.Keypair)
			if !ok {
				return fmt.Errorf("ssv keypair err")
			}

			// load seed from keystore
			seed, err := loadSeed(cfg.KeystorePath)
			if err != nil {
				return err
			}

			t, err := task.NewTask(cfg, seed, kp, ssvkp)
			if err != nil {
				return err
			}
			logrus.Info("starting services...")
			err = t.Start()
			if err != nil {
				logrus.Errorf("start err: %s", err)
				return err
			}
			defer func() {
				logrus.Infof("shutting down task ...")
				t.Stop()
			}()

			// <-ctx.Done()
			return nil
		},
	}

	cmd.Flags().String(flagConfigPath, defaultConfigPath, "Config file path")
	cmd.Flags().String(flagOperatorsPath, defaultOperatorsPath, "Operators file path")
	cmd.Flags().String(flagLogLevel, logrus.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")

	return cmd
}
