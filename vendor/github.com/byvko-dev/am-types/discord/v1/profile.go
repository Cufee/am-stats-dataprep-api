package discord

// Full Discord Profile and auth token data per https://discord.com/developers/docs/resources/user#user-object
type DiscordProfile struct {
	Token AccessToken `json:"token" firestore:"token"`

	ID          string `json:"id" firestore:"id"`
	Username    string `json:"username" firestore:"username"`
	Discrim     string `json:"discriminator" firestore:"discriminator"`
	Avatar      string `json:"avatar" firestore:"avatar"`
	Mfa         bool   `json:"mfa_enabled" firestore:"mfa_enabled"`
	Banner      string `json:"banner" firestore:"banner"`
	AccentColor int    `json:"accent_color" firestore:"accent_color"`
	Locale      string `json:"locale" firestore:"locale"`
	Verified    bool   `json:"verified" firestore:"verified"`
	Email       string `json:"email" firestore:"email"`
	Flags       int    `json:"flags" firestore:"flags"`
	PremiumType int    `json:"premium_type" firestore:"premium_type"`
	PublicFlags int    `json:"public_flags" firestore:"public_flags"`
}
