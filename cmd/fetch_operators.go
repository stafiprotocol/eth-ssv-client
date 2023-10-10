package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

type SsvOperators struct {
	Operators []Operator `toml:"operators"`
}
type Operator struct {
	Id     uint64 `toml:"id"`
	Pubkey string `toml:"pubkey"`
}

const networkMainnet = "mainnet"
const networkPrater = "prater"

var supportedNetwork = map[string]bool{
	networkMainnet: true,
	networkPrater:  true,
}

func fetchOperatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch-operators",
		Args:  cobra.ExactArgs(0),
		Short: fmt.Sprintf("Fetch operators pubkey from ssv api and save to file %s", defaultOperatorsPath),
		RunE: func(cmd *cobra.Command, args []string) error {
			network, err := cmd.Flags().GetString(flagNetwork)
			if err != nil {
				return err
			}
			if !supportedNetwork[network] {
				return fmt.Errorf("unsupported network: %s", network)
			}

			operators, err := utils.GetSsvOperators(network)
			if err != nil {
				return err
			}

			ssvOperators := SsvOperators{Operators: make([]Operator, len(operators))}
			for i, op := range operators {
				_, err := base64.StdEncoding.DecodeString(op.PublicKey)
				if err != nil {
					fmt.Printf("operator: %d pubkey decode err, will skip this operator. pubkey: %s\n", op.ID, op.PublicKey)
					continue
				}
				ssvOperators.Operators[i] = Operator{Id: uint64(op.ID), Pubkey: op.PublicKey}
			}

			filePath := defaultOperatorsPath
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return errors.Errorf("failed to open file %s: %s", filePath, err)
			}
			defer f.Close()

			err = toml.NewEncoder(f).Encode(ssvOperators)
			if err != nil {
				return err
			}
			fmt.Printf("fetch %d operators and save to %s ok\n", len(ssvOperators.Operators), defaultOperatorsPath)
			return nil
		},
	}

	cmd.Flags().String(flagNetwork, defaultNetwork, fmt.Sprintf("Network (%s|%s)", networkMainnet, networkPrater))

	return cmd
}
