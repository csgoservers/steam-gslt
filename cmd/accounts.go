package cmd

import (
	"fmt"

	"github.com/csgoservers/steam-gameserver-service/pkg/client"
	"github.com/spf13/cobra"
)

var getAccountListCmd = &cobra.Command{
	Use:   "GetAccountList",
	Short: "Gets a list of game server accounts with their logon tokens",
	Run: func(cmd *cobra.Command, args []string) {
		// check if key flag is set
		err := checkRequiredFlags(cmd)
		if err != nil {
			rootCmd.Help()
			return
		}
		executeCmd(cmd.Flag("key").Value.String())
	},
}

func executeCmd(key string) {
	steam := client.New(key)
	accounts, err := steam.GetAccountList()
	if err != nil {
		fmt.Printf("Error getting accounts: %v\n", err)
		return
	}
	for _, server := range accounts.Servers {
		fmt.Printf("Steam ID: %s\tToken: %s\n", server.SteamID, server.LoginToken)
	}
}
