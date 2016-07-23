package types

type Request struct {
	FbToken string `json:"fb_token"`
	UserId string `json:"user_id"`
	Groups []string `json:"groups"`
	IgnoreAlbums []string `json:"ignore_albums"`
	Keywords []string `json:"keywords"`
}
