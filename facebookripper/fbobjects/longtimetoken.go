package objects

type LongTimeToken struct {
	ObtainDate int `json:"obtain_date"`
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}
