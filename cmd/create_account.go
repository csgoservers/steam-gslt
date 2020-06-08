package main

import (
	"fmt"

	"github.com/csgoservers/steam-gameserver-service/pkg/client"
	"github.com/spf13/cobra"
)

var createAccountCmd = &cobra.Command{
	Use:   "CreateAccount",
	Short: "Creates a persistent game server account",
	Run: func(cmd *cobra.Command, args []string) {
		// check if key flag is set
		err := checkRequiredFlags()
		if err != nil {
			rootCmd.Help()
			return
		}
		if appID == 0 {
			cmd.Help()
			return
		}
		executeCreateAccountCmd()
	},
}

func executeCreateAccountCmd() {
	steam := client.New(key)
	account, err := steam.CreateAccount(appID, memo)
	if err != nil {
		fmt.Printf("Error getting accounts: %v\n", err)
		return
	}
	fmt.Println(account.LoginToken)
}
