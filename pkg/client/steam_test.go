package client

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeServer(fn http.HandlerFunc) (SteamService, func()) {
	server := httptest.NewServer(fn)
	service := newService("abc")
	service.url = server.URL
	steam := SteamService{}
	steam.service = service
	return steam, server.Close
}

func TestGetAccountList(t *testing.T) {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.ReadFile("../../testdata/fixture_account_list.json")
		if err != nil {
			assert.NoError(t, err)
		}
		w.Write(f)
	})
	steam, close := fakeServer(fn)
	defer close()

	accounts, err := steam.GetAccountList()
	assert.NoError(t, err)
	assert.Equal(t, "76561197900265728", accounts.Actor)
	assert.True(t, accounts.IsBanned)
	assert.Equal(t, 1, len(accounts.Servers))

	account := accounts.Servers[0]
	assert.Equal(t, "1111", account.SteamID)
	assert.Equal(t, "abc", account.LoginToken)
	assert.Equal(t, "test", account.Memo)
	assert.Equal(t, 730, account.AppID)
	assert.Equal(t, 1589448113, account.LastLogon)
	assert.True(t, account.IsExpired)
	assert.False(t, account.IsDeleted)
}

func TestCreateAccount(t *testing.T) {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.ReadFile("../../testdata/fixture_create_account.json")
		if err != nil {
			assert.NoError(t, err)
		}
		w.Write(f)
	})
	steam, close := fakeServer(fn)
	defer close()

	account, err := steam.CreateAccount(730, "hello world")
	assert.NoError(t, err)
	assert.Equal(t, 730, account.AppID)
	assert.Equal(t, "hello world", account.Memo)
	assert.Equal(t, "80068392925402169", account.SteamID)
	assert.Equal(t, "D212EAB4B33A0005CA4CD483AAAA4C9E", account.LoginToken)
}

func TestSetMemo(t *testing.T) {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.ReadFile("../../testdata/fixture_empty_response.json")
		if err != nil {
			assert.NoError(t, err)
		}
		w.Write(f)
	})
	steam, close := fakeServer(fn)
	defer close()

	err := steam.SetMemo(80068392925402169, "hello world")
	assert.NoError(t, err)
}
