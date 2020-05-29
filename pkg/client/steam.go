package client

import (
	"encoding/json"
	"net/url"
	"strconv"

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

// CreateAccount creates a persistent game server account
func (s *SteamService) CreateAccount(appID uint32, memo string) (*api.Account, error) {
	data := url.Values{}
	data.Add("appid", strconv.Itoa(int(appID)))
	data.Add("memo", memo)

	// TODO: extract return value to convert to Account
	_, err := s.service.post("CreateAccount", data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
