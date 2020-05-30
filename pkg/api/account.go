package api

// ServerToken is the Steam info for one steamid
type ServerToken struct {
	SteamID    string `json:"steamid"`
	LoginToken string `json:"login_token"`
	Memo       string `json:"memo"`
	IsDeleted  bool   `json:"is_deleted"`
	IsExpired  bool   `json:"is_expired"`
	LastLogon  int    `json:"rt_last_logon"`
	AppID      int    `json:"appid"`
}

// Account is the response that contains data about all server tokens
// emited for the current account.
type Account struct {
	Actor          string        `json:"actor"`
	IsBanned       bool          `json:"is_banned"`
	Expires        int           `json:"expires"`
	LastActionTime int           `json:"last_action_time"`
	Servers        []ServerToken `json:"servers"`
}
