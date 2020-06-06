package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	// key for the Steam web API
	key string

	rootCmd = &cobra.Command{
		Use:   "steamgs-cli",
		Short: "Command line application to request Steam game server services",
	}
)

// Execute the root command for steamgs-cli
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Steam web API key (required)")
	rootCmd.AddCommand(getAccountListCmd)
}

func checkRequiredFlags(cmd *cobra.Command) error {
	key := cmd.Flag("key").Value.String()
	if key == "" {
		return errors.New("steam API key is required. See help for more info")
	}
	return nil
}
