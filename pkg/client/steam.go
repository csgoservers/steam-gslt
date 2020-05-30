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
func (s *SteamService) GetAccountList() (*api.Account, error) {
	result, err := s.service.get("GetAccountList")
	if err != nil {
		return nil, err
	}
	var wrapper api.Account
	err = json.Unmarshal(result, &wrapper)
	if err != nil {
		return nil, err
	}
	return &wrapper, nil
}

// CreateAccount creates a persistent game server token
func (s *SteamService) CreateAccount(appID int, memo string) (*api.ServerToken, error) {
	data := url.Values{}
	data.Add("appid", strconv.Itoa(appID))
	data.Add("memo", memo)

	result, err := s.service.post("CreateAccount", data)
	if err != nil {
		return nil, err
	}
	var token api.ServerToken
	err = json.Unmarshal(result, &token)
	if err != nil {
		return nil, err
	}
	// fill object with data that is not present
	// in the response object
	token.AppID = appID
	token.Memo = memo
	return &token, nil
}
