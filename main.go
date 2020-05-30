package main

import (
	"flag"
	"log"

	"github.com/csgoservers/steam-gameserver-service/pkg/client"
)

func main() {
	key := flag.String("key", "", "Steam API key")
	flag.Parse()

	if *key == "" {
		log.Println("Steam API key can't be empty")
		return
	}
	steam := client.New(*key)
	response, err := steam.GetAccountList()
	if err != nil {
		log.Printf("Error calling GetAccountList: %v", err)
		return
	}
	log.Println(response.Servers)
}
