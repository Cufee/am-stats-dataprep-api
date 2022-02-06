package types

// Request to get player stats
type StatsRequest struct {
	OptionsProfile string `json:"options_profile"`
	PlayerID       string `json:"player_id"`
	Locale         string `json:"locale"`
}
