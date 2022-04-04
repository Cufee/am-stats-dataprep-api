package types

import "github.com/byvko-dev/am-types/dataprep/v1/block"

type StatsResponse struct {
	StatusIcons []block.Block `json:"statusIcons" bson:"statusIcons"`
	Cards       []block.Block `json:"cards" bson:"cards"`
	FailedCards []string      `json:"failedCards" bson:"failedCards"`
	LastBattle  int           `json:"lastBattle" bson:"lastBattle"`
	Style       string        `json:"style" bson:"style"`
}
