package dataprep

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

func ShotAccuracyBlock(layoutName string, session, allTime statistics.StatsFrame) block.Block {
	layout := layouts.Load(layoutName, shared.Accuracy)
	values := make(logic.Values)
	values[logic.SessionValue] = float64(session.Hits)
	values[logic.SessionOf] = float64(session.Shots)
	values[logic.AllTimeOf] = float64(allTime.Shots)
	values[logic.AllTimeValue] = float64(allTime.Hits)
	layout.Values = values
	return layout.ToBlock()
}
