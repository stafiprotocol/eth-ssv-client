package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var (
	appName = "ssv-client"
)

const (
	flagKeystorePath  = "keystore_path"
	flagLogLevel      = "log_level"
	flagConfigPath    = "config"
	flagOperatorsPath = "operators"
	flagViewMode      = "view_mode"
	flagNetwork       = "network"

	defaultKeystorePath  = "./keys"
	defaultConfigPath    = "./config.toml"
	defaultOperatorsPath = "./operators.toml"
	defaultNetwork       = "prater"
)

// NewRootCmd returns the root command.
func NewRootCmd() *cobra.Command {
	// RootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   appName,
		Short: "ssv-client",
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {

		return nil
	}

	rootCmd.AddCommand(
		importAccountCmd(),
		importMnemonicCmd(),
		startSsvCmd(),
		fetchOperatorsCmd(),
		versionCmd(),
	)
	return rootCmd
}

func Execute() {
	cobra.EnableCommandSorting = false

	rootCmd := NewRootCmd()
	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	ctx := context.Background()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
