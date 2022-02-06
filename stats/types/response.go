package types

type StatsResponse struct {
	StatusIcons  []StatsBlockRowContent `json:"statusIcons"`
	Cards        []StatsCard            `json:"cards"`
	FailedCards  []string               `json:"failedCards"`
	StyleProfile string                 `json:"styleProfile"`
}
