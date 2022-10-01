package stats

// ExportData - Struct to export final data for use in Python bot
type PlayerRawStats struct {
	PlayerCache   PlayerProfile `json:"player_cache"`
	PlayerDetails PlayerDetails `json:"player_details"`
	SessionStats  Session       `json:"session"`
	LastSession   RetroSession  `json:"last_session"`
	Error         string        `json:"error"`
}
