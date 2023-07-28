VERSION := $(shell git describe --tags)
COMMIT  := $(shell git log -1 --format='%H')

all: build

LD_FLAGS = -X github.com/stafiprotocol/eth-ssv-client/cmd.Version=$(VERSION) \
	-X github.com/stafiprotocol/eth-ssv-client/cmd.Commit=$(COMMIT) \

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

get:
	@echo "  >  \033[32mDownloading & Installing all the modules...\033[0m "
	go mod tidy && go mod download

build:
	@echo " > \033[32mBuilding ssv-client...\033[0m "
	go build -mod readonly $(BUILD_FLAGS) -o build/ssv-client main.go

build-linux:
	@GOOS=linux GOARCH=amd64 go build --mod readonly $(BUILD_FLAGS) -o ./build/ssv-client main.go

install:
	@echo " > \033[32mInstalling ssv-client...\033[0m "
	go install -mod readonly $(BUILD_FLAGS) ./...

abi:
	@echo " > \033[32mGenabi...\033[0m "
	abigen --abi ./bindings/NodeDeposit/nodedeposit_abi.json --pkg node_deposit --type NodeDeposit --out ./bindings/NodeDeposit/NodeDeposit.go
	abigen --abi ./bindings/LightNode/lightnode_abi.json --pkg light_node --type LightNode --out ./bindings/LightNode/LightNode.go
	abigen --abi ./bindings/SuperNode/supernode_abi.json --pkg super_node --type SuperNode --out ./bindings/SuperNode/SuperNode.go
	abigen --abi ./bindings/StakingPool/stakingpool_abi.json --pkg staking_pool --type StakingPool --out ./bindings/StakingPool/StakingPool.go
	abigen --abi ./bindings/Settings/networksettings_abi.json --pkg network_settings --type NetworkSettings --out ./bindings/Settings/NetworkSettings.go
	abigen --abi ./bindings/NetworkBalances/networkbalances_abi.json --pkg network_balances --type NetworkBalances --out ./bindings/NetworkBalances/NetworkBalances.go
	abigen --abi ./bindings/Reth/reth_abi.json --pkg reth --type Reth --out ./bindings/Reth/Reth.go
	abigen --abi ./bindings/UserDeposit/userdeposit_abi.json --pkg user_deposit --type UserDeposit --out ./bindings/UserDeposit/UserDeposit.go
	abigen --abi ./bindings/Storage/storage_abi.json --pkg storage --type Storage --out ./bindings/Storage/Storage.go
	abigen --abi ./bindings/NodeManager/nodemanager_abi.json --pkg node_manager --type NodeManager --out ./bindings/NodeManager/NodeManager.go
	abigen --abi ./bindings/StakingPoolManager/stakingpoolmanager_abi.json --pkg staking_pool_manager --type StakingPoolManger --out ./bindings/StakingPoolManager/StakingPoolManager.go
	abigen --abi ./bindings/Distributor/distributor_abi.json --pkg distributor --type Distributor --out ./bindings/Distributor/Distributor.go
	abigen --abi ./bindings/FeePool/feepool_abi.json --pkg fee_pool --type FeePool --out ./bindings/FeePool/FeePool.go
	abigen --abi ./bindings/SuperNodeFeePool/supernodefeepool_abi.json --pkg super_node_fee_pool --type SuperNodeFeePool --out ./bindings/SuperNodeFeePool/SuperNodeFeePool.go
	abigen --abi ./bindings/Withdraw/withdraw_abi.json --pkg withdraw --type Withdraw --out ./bindings/Withdraw/Withdraw.go
	abigen --abi ./bindings/StafiEther/stafiether_abi.json --pkg stafi_ether --type StafiEther --out ./bindings/StafiEther/StafiEther.go
	abigen --abi ./bindings/StakePortalRate/stakeportalrate_abi.json --pkg stake_portal_rate --type StakePortalRate --out ./bindings/StakePortalRate/StakePortalRate.go
	abigen --abi ./bindings/SsvNetwork/ssvnetwork_abi.json --pkg ssv_network --type SsvNetwork --out ./bindings/SsvNetwork/SsvNetwork.go
	abigen --abi ./bindings/SsvNetworkViews/ssvnetworkviews_abi.json --pkg ssv_network_views --type SsvNetworkViews --out ./bindings/SsvNetworkViews/SsvNetworkViews.go
	abigen --abi ./bindings/SsvClusters/ssvclusters_abi.json --pkg ssv_clusters --type SsvClusters --out ./bindings/SsvClusters/SsvClusters.go
	abigen --abi ./bindings/Erc20/erc20_abi.json --pkg erc20 --type Erc20 --out ./bindings/Erc20/Erc20.go


clean:
	@echo " > \033[32mCleanning build files ...\033[0m "
	rm -rf build
fmt :
	@echo " > \033[32mFormatting go files ...\033[0m "
	go fmt ./...

swagger:
	@echo "  >  \033[32mBuilding swagger docs...\033[0m "
	swag init --parseDependency

get-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest

lint:
	golangci-lint run ./... --skip-files ".+_test.go"

.PHONY: all lint test race msan tools clean build
