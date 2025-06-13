package cmd

import (
	"encoding/hex"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	"github.com/stafiprotocol/chainbridge/utils/keystore"
)

func exportAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-account",
		Args:  cobra.ExactArgs(0),
		Short: "Export account",
		RunE: func(cmd *cobra.Command, args []string) error {
			keystorePath, err := cmd.Flags().GetString(flagKeystorePath)
			if err != nil {
				return err
			}
			fmt.Printf("keystore path: %s\n", keystorePath)
			accountStr, err := cmd.Flags().GetString(flagAccount)
			if err != nil {
				return err
			}

			// load ssv account
			ssvkpI, err := keystore.KeypairFromAddress(accountStr, keystore.EthChain, keystorePath, false)
			if err != nil {
				return err
			}
			ssvkp, ok := ssvkpI.(*secp256k1.Keypair)
			if !ok {
				return fmt.Errorf("ssv keypair err")
			}

			fmt.Printf("privateKey: %s\n", hex.EncodeToString(ethCrypto.FromECDSA(ssvkp.PrivateKey())))
			return nil
		},
	}
	cmd.Flags().String(flagKeystorePath, defaultKeystorePath, "Keystore file path")
	cmd.Flags().String(flagAccount, "", "Account hex address")
	return cmd
}
