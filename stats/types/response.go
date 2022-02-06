package types

type StatsResponse struct {
	StatusIcons []StatsBlockRowContent `json:"status_icons"`
	Cards       []StatsCard            `json:"cards"`
	FailedCards []string               `json:"failed_cards"`
}
