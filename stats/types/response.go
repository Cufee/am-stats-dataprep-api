package types

type StatsResponse struct {
	StatusIcons []StatsBlockRowContent `json:"statusIcons" bson:"statusIcons"`
	Cards       []StatsCard            `json:"cards" bson:"cards"`
	FailedCards []string               `json:"failedCards" bson:"failedCards"`
	StylePreset string                 `json:"stylePreset" bson:"stylePreset"`
	LastBattle  int                    `json:"lastBattle" bson:"lastBattle"`
}
