// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ListenAddr              string
	Eth1Endpoint            string // url for eth1 rpc endpoint
	Eth2Endpoint            string // url for eth2 rpc endpoint
	ArbitrumEndpoint        string // url for arbitrum rpc endpoint
	LogFilePath             string
	From                    string   // address of voter
	SuperNodeAccount        string   // ssv use
	SsvAccount              string   // ssv use
	Operators               []uint64 // ssv use
	ClusterInitSsvAmount    string   //ssv use
	KeystorePath            string
	GasLimit                string
	MaxGasPrice             string
	UnstakingStartTimestamp uint64 // used by api server
	RunClientStartTimestamp uint64 // used by api server
	StafiInfoEndpoint       string // used by api server

	EraCount uint64 // used for statistic

	Contracts Contracts

	Db Db
}

type Contracts struct {
	StorageContractAddress         string
	ArbitrumStakePortalRateAddress string
	SsvNetworkAddress              string
	SsvNetworkViewsAddress         string
	SsvTokenAddress                string
}

type Db struct {
	Host string
	Port string
	User string
	Pwd  string
	Name string
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
