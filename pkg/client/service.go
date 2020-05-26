package service

import (
	"net/http"

	"github.com/csgoservers/steam-gameserver-service/pkg/api"
)

// GameServerService is the base structure to call methods of the
// Steam IGameServersService interface.
type GameServerService struct {
	http *http.Client
}

// New creates a new service client to make requests to the Steam API
func New() *GameServerService {
	client := &http.Client{}
	return &GameServerService{client}
}

// GetAccountList returns a list of game server accounts with their
// connection tokens.
func (g *GameServerService) GetAccountList() (*api.Account, error) {
	return nil, nil
}
