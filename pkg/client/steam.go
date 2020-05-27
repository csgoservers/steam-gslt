package client

import (
	"encoding/json"

	"github.com/csgoservers/steam-gameserver-service/pkg/api"
)

// SteamService to retrieve and send data
type SteamService struct {
	service *gameServerService
}

// New creates a SteamService to make network requests to the Steam API
func New(apiKey string) *SteamService {
	return &SteamService{newService(apiKey)}
}

// GetAccountList returns a list of game server accounts with their
// connection tokens.
func (s *SteamService) GetAccountList() (*api.AccountResponse, error) {
	result, err := s.service.get("GetAccountList")
	if err != nil {
		return nil, err
	}
	var wrapper api.AccountWrapperResponse
	err = json.Unmarshal(result, &wrapper)
	if err != nil {
		return nil, err
	}
	return &wrapper.Raw, nil
}
