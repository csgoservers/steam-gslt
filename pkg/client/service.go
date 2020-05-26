package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// gameServerService is the base structure to call methods of the
// Steam IGameServersService interface.
type gameServerService struct {
	http *http.Client
	url  string
}

// New creates a new service client to make requests to the Steam API
func newService() *gameServerService {
	client := &http.Client{}
	return &gameServerService{client, "https://api.steampowered.com/IGameServersService/"}
}

// get executes a request to retrieve some data from the service
func (g *gameServerService) get(method string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/%s", g.url, method, "v1")

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
