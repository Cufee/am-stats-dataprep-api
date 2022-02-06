package types

// Request to get player stats
type StatsRequest struct {
	OptionsProfile string `json:"options_profile"`
	SettingsID     string `json:"settings_id"`
	PlayerID       int    `json:"player_id"`
	Profile        string `json:"profile"`
	Locale         string `json:"locale"`
	Realm          string `json:"realm"`
}
