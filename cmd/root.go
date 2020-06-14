package main

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	// key for the Steam web API
	key   string
	memo  string
	appID int

	rootCmd = &cobra.Command{
		Use:   "gslt-cli",
		Short: "Command line application to request Steam game server services",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Steam web API key (required)")

	createAccountCmd.PersistentFlags().IntVarP(&appID, "appid", "a", 0, "The app to use the account for")
	createAccountCmd.PersistentFlags().StringVarP(&memo, "memo", "m", "", "The memo to set on the new account")

	rootCmd.AddCommand(getAccountListCmd)
	rootCmd.AddCommand(createAccountCmd)
}

func checkRequiredFlags() error {
	if key == "" {
		return errors.New("steam API key is required. See help for more info")
	}
	return nil
}
