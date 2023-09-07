# eth-ssv-client

Client for StaFi staking pool and SSV

## Build

* [Go 1.20+](https://golang.org/dl/), for example:

```bash
cd $HOME
wget -O go1.20.3.linux-amd64.tar.gz https://go.dev/dl/go1.20.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz && rm go1.20.3.linux-amd64.tar.gz
echo 'export GOROOT=/usr/local/go' >> $HOME/.bashrc
echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
echo 'export GO111MODULE=on' >> $HOME/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> $HOME/.bashrc && . $HOME/.bashrc
go version
```

* make 

```base
 git clone https://github.com/stafiprotocol/eth-ssv-client.git
 cd eth-ssv-client
 make
```

## Usage

```
./build/ssv-client

ssv-client

Usage:
  ssv-client [command]

Available Commands:
  import-account      Import account
  import-val-mnemonic Import mnemonic of validators
  start               Start ssv client
  version             Show version information
  help                Help about any command

Flags:
  -h, --help   help for ssv-client

Use "ssv-client [command] --help" for more information about a command.
```

## Feature

Automatically:
* Generate valdiator keys from mnemonic
* Deposit/Stake on StaFi staking pool
* ~~Select/Monitor operators from SSV api~~ (Due to ssv api not support on mainnet)
* Onboard/Offboard validators on SSV
* Deposit/Withdraw/Reactive cluster on SSV
* Hold StaFi staking pool ejector service