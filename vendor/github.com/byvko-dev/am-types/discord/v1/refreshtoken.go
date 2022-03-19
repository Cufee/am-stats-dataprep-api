package discord

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type" firestore:"grant_type"`
	ClientSecret string `json:"client_secret" firestore:"client_secret"`
	ClientID     string `json:"client_id" firestore:"client_id"`
	RefreshToken string `json:"refresh_token" firestore:"refresh_token"`
}
