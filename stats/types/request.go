package types

// Request to get player stats
type BasicStatsRequest struct {
	OptionsProfile string `json:"options_profile"`
	PlayerID       int    `json:"player_id"`
	Profile        string `json:"profile"`
	Locale         string `json:"locale"`
	Realm          string `json:"realm"`
}
