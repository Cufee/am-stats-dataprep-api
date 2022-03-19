package discord

type AccessTokenResponse struct {
	Error string `json:"error,omitempty"`
	AccessToken
}

type AccessToken struct {
	RefreshToken string `json:"refresh_token" firestore:"refresh_token"`
	Token        string `json:"access_token" firestore:"access_token"`
	Type         string `json:"token_type" firestore:"token_type"`
	Exp          int    `json:"expires_in" firestore:"expires_in"`
	Scope        string `json:"scope" firestore:"scope"`
}

type TokenRequest struct {
	Code         string `json:"code" firestore:"code"`
	ClientID     string `json:"client_id" firestore:"client_id"`
	GrantType    string `json:"grant_type" firestore:"grant_type"`
	RedirectURI  string `json:"redirect_uri" firestore:"redirect_uri"`
	ClientSecret string `json:"client_secret" firestore:"client_secret"`
}
