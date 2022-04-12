package generators

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

func BlockFromStats(layoutName string, definition logic.Definition, session, allTime statistics.StatsFrame, printer func(string) string) *block.Block {
	layout := layouts.LoadDefinition(layoutName, definition)
	values := make(logic.Values)

	switch definition.ValueKind {
	case logic.BattlesOverOne:
		values[logic.SessionValue] = float64(session.Battles)
		values[logic.SessionOf] = float64(1)
		values[logic.AllTimeOf] = float64(1)
		values[logic.AllTimeValue] = float64(allTime.Battles)

	case logic.WinsOverBattles:
		values[logic.SessionValue] = float64(session.Wins)
		values[logic.SessionOf] = float64(session.Battles)
		values[logic.AllTimeOf] = float64(allTime.Battles)
		values[logic.AllTimeValue] = float64(allTime.Wins)

	case logic.DamageOverBattles:
		values[logic.SessionValue] = float64(session.DamageDealt)
		values[logic.SessionOf] = float64(session.Battles)
		values[logic.AllTimeOf] = float64(allTime.Battles)
		values[logic.AllTimeValue] = float64(allTime.DamageDealt)

	case logic.HitsOverShots:
		values[logic.SessionValue] = float64(session.Hits)
		values[logic.SessionOf] = float64(session.Shots)
		values[logic.AllTimeOf] = float64(allTime.Shots)
		values[logic.AllTimeValue] = float64(allTime.Hits)

	default:
		return nil
	}

	layout.Values = values
	return layout.ToBlock(printer)
}
