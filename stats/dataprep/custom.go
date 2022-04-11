package dataprep

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func CustomBlock(layout *logic.Layout, sessionValue, sessionOf, allTimeValue, allTimeOf float64) block.Block {
	values := make(logic.Values)
	values[logic.SessionValue] = sessionValue
	values[logic.SessionOf] = sessionOf
	values[logic.AllTimeOf] = allTimeValue
	values[logic.AllTimeValue] = allTimeOf
	layout.Values = values
	return layout.ToBlock()
}
