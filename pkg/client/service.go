package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// gameServerService is the base structure to call methods of the
// Steam IGameServersService interface.
type gameServerService struct {
	http *http.Client
	url  string
	key  string
}

// New creates a new service client to make requests to the Steam API
func newService(key string) *gameServerService {
	client := &http.Client{}
	return &gameServerService{client, "https://api.steampowered.com/IGameServersService", key}
}

// get executes a request to retrieve some data from the service
func (g *gameServerService) get(method string) ([]byte, error) {
	url := g.buildURL(method)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := g.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// post executes a requests to change data using the Steam service
func (g *gameServerService) post(method string, data url.Values) ([]byte, error) {
	url := g.buildURL(method)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := g.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (g *gameServerService) buildURL(method string) string {
	return fmt.Sprintf("%s/%s/%s?key=%s", g.url, method, "v1", g.key)
}
