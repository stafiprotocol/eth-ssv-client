// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Eth1Endpoint     string // url for eth1 rpc endpoint
	Eth2Endpoint     string // url for eth2 rpc endpoint
	LogFilePath      string
	SuperNodeAccount string     // ssv use
	SsvAccount       string     // ssv use
	Clusters         [][]uint64 // ssv use
	KeystorePath     string
	GasLimit         string
	MaxGasPrice      string

	Contracts Contracts
}

type Contracts struct {
	StorageContractAddress         string
	ArbitrumStakePortalRateAddress string
	SsvNetworkAddress              string
	SsvNetworkViewsAddress         string
	SsvTokenAddress                string
}

func Load(configFilePath string) (*Config, error) {
	var cfg = Config{}
	if err := loadSysConfig(configFilePath, &cfg); err != nil {
		return nil, err
	}
	if len(cfg.LogFilePath) == 0 {
		cfg.LogFilePath = "./log_data"
	}

	return &cfg, nil
}

func loadSysConfig(path string, config *Config) error {
	_, err := os.Open(path)
	if err != nil {
		return err
	}
	if _, err := toml.DecodeFile(path, config); err != nil {
		return err
	}
	fmt.Println("load config success")
	return nil
}
