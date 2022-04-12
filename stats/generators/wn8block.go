package generators

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func WN8BlockFromStats(layoutName string, definition logic.Definition, session, allTime int, printer func(string) string) *block.Block {
	layout := layouts.LoadDefinition(layoutName, definition)
	values := make(logic.Values)

	values[logic.SessionValue] = float64(session)
	values[logic.SessionOf] = float64(1)
	values[logic.AllTimeOf] = float64(1)
	values[logic.AllTimeValue] = float64(allTime)

	layout.Values = values
	return layout.ToBlock(printer)
}
