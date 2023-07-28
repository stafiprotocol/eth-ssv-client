package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/chainbridge/utils/keystore"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/mnemonic"
	"github.com/tyler-smith/go-bip39"
)

const ssvValidatorKeyFileName = "ssv_validator.key"

func importMnemonicCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-val-mnemonic",
		Args:  cobra.ExactArgs(0),
		Short: "Import mnemonic for ssv",
		RunE: func(cmd *cobra.Command, args []string) error {
			keystorePath, err := cmd.Flags().GetString(flagKeystorePath)
			if err != nil {
				return err
			}
			fmt.Printf("keystore path: %s\n", keystorePath)
			logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
			if err != nil {
				return err
			}
			logLevel, err := logrus.ParseLevel(logLevelStr)
			if err != nil {
				return err
			}
			logrus.SetLevel(logLevel)

			return importSeedFromMnemonic(keystorePath)
		},
	}
	cmd.Flags().String(flagKeystorePath, defaultKeystorePath, "Keystore file path")
	cmd.Flags().String(flagLogLevel, logrus.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")
	return cmd
}

func importSeedFromMnemonic(keypath string) error {
	var err error

	mnemonicBts := keystore.GetPassword("Enter mnemonic:")
	mnemonicStr := string(mnemonicBts)
	if !bip39.IsMnemonicValid(mnemonicStr) {
		return fmt.Errorf("not valid mnemonic")
	}
	seed := mnemonic.NewSeed(mnemonicStr, "")

	fp, err := filepath.Abs(keypath + "/" + ssvValidatorKeyFileName)
	if err != nil {
		return fmt.Errorf("invalid filepath: %s", err)
	}

	if _, err := os.Stat(fp); err != nil {
		err := os.MkdirAll(filepath.Dir(fp), 0700)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("key for ssv already exist, please remove old file before import")
	}

	password := keystore.GetPassword("password for key:")

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			logrus.Error("generate keypair: could not close keystore file")
		}
	}()

	err = encryptAndWriteToFile(file, seed, password)
	if err != nil {
		return fmt.Errorf("could not write key to file: %s", err)
	}
	logrus.WithFields(logrus.Fields{
		"file": fp,
	}).Info("mnemonic generated")

	return nil
}

func encryptAndWriteToFile(file *os.File, seed []byte, password []byte) error {
	ciphertext, err := keystore.Encrypt(seed, password)
	if err != nil {
		return err
	}

	_, err = file.Write(ciphertext)
	return err
}

func loadSeed(keypath string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s", keypath, ssvValidatorKeyFileName)
	// Make sure key exists before prompting password
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("key file not found: %s", path)
	}

	pswd := keystore.GetPassword(fmt.Sprintf("Enter password for key %s:", path))

	return ReadFromFileAndDecrypt(path, pswd)
}

// ReadFromFileAndDecrypt reads ciphertext from a file and decrypts it using the password into a `crypto.PrivateKey`
func ReadFromFileAndDecrypt(filename string, password []byte) ([]byte, error) {
	fp, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filepath.Clean(fp))
	if err != nil {
		return nil, err
	}

	return keystore.Decrypt(data, password)
}
