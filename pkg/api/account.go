package api

// Account is the Steam info for one steamid
type Account struct {
	SteamID    string `json:"steamid"`
	LoginToken string `json:"login_token"`
	Memo       string `json:"memo"`
	IsDeleted  bool   `json:"is_deleted"`
	IsExpired  bool   `json:"is_expired"`
	LastLogon  int    `json:"rt_last_logon"`
	AppID      int    `json:"appid"`
}

// AccountResponse is the response that contains data about Accounts
type AccountResponse struct {
	Actor          string    `json:"actor"`
	IsBanned       bool      `json:"is_banned"`
	Expires        int       `json:"expires"`
	LastActionTime int       `json:"last_action_time"`
	Accounts       []Account `json:"servers"`
}

// AccountWrapperResponse is a wrapper around AccountResponse
type AccountWrapperResponse struct {
	Raw AccountResponse `json:"response"`
}
