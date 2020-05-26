package client

import "github.com/csgoservers/steam-gameserver-service/pkg/api"

// SteamService to retrieve and send data
type SteamService struct {
	service *gameServerService
}

// New creates a SteamService to make network requests to the Steam API
func New() *SteamService {
	return &SteamService{newService()}
}

// GetAccountList returns a list of game server accounts with their
// connection tokens.
func (s *SteamService) GetAccountList() (*[]api.Account, error) {
	_, err := s.service.get("GetAccountList")
	return nil, err
}
