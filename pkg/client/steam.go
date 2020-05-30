package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	"github.com/csgoservers/steam-gameserver-service/pkg/api"
)

var emptyJSON = []byte(`{}`)

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

// SetMemo changes the note associated to the game server. This memo only serves as
// a reminder of the purpose of the server token.
func (s *SteamService) SetMemo(steamID int, memo string) error {
	data := url.Values{}
	data.Add("steamid", strconv.Itoa(steamID))
	data.Add("memo", memo)

	_, err := s.service.post("SetMemo", data)
	return err
}

// ResetLoginToken generates a new token for the current game server.
func (s *SteamService) ResetLoginToken(steamID int) (*api.ServerToken, error) {
	data := url.Values{}
	data.Add("steamid", strconv.Itoa(steamID))

	result, err := s.service.post("ResetLoginToken", data)
	if err != nil {
		return nil, err
	}
	var token api.ServerToken
	err = json.Unmarshal(result, &token)
	if err != nil {
		return nil, err
	}
	// fill steam ID on the returned server token
	token.SteamID = strconv.Itoa(steamID)
	return &token, nil
}

// DeleteAccount deletes the given Steam ID.
func (s *SteamService) DeleteAccount(steamID int) error {
	data := url.Values{}
	data.Add("steamid", strconv.Itoa(steamID))

	_, err := s.service.post("DeleteAccount", data)
	return err
}

// QueryLoginToken generates a new token for the current game server.
func (s *SteamService) QueryLoginToken(loginToken string) (*api.Account, error) {
	data := url.Values{}
	data.Add("login_token", loginToken)

	result, err := s.service.post("QueryLoginToken", data)
	if err != nil {
		return nil, err
	}
	// if result is empty, then the game server token is not valid and we return
	// an error. To detect if result is not valid we check if the array size is
	// equals to 0 or if the result is equal to an empty JSON.
	if len(result) == 0 || bytes.Compare(emptyJSON, result) == 0 {
		return nil, errors.New("game server login token is not valid")
	}
	var account api.Account
	err = json.Unmarshal(result, &account)
	if err != nil {
		return nil, err
	}
	// fill account with the steam id and login token
	var tokens []api.ServerToken
	token := api.ServerToken{}
	token.LoginToken = loginToken

	// bacause steamid is present in the response
	// we need to parse it manually to be able to set this ID
	// in the ServerToken object.
	var raw map[string]json.RawMessage
	_ = json.Unmarshal(result, &raw)
	_ = json.Unmarshal(raw["steamid"], &token.SteamID)

	tokens = append(tokens, token)
	account.Servers = tokens
	return &account, nil
}
