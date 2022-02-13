package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"

	stats "byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func AvarageDamageBlock(input types.DataprepInput) (stats.StatsBlock, error) {
	if input.Stats.Session.Battles == 0 {
		return stats.StatsBlock{}, fmt.Errorf("session stats have 0 battles")
	}

	var block stats.StatsBlock
	block.Tags = append(block.Tags, input.Options.Block.GenerationTag+"Block")
	block.Rows = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), (input.Stats.Session.Battles), (input.Stats.AllTime.DamageDealt), (input.Stats.AllTime.Battles))
	return block, nil
}

func DamageDoneBlock(input types.DataprepInput) (stats.StatsBlock, error) {
	if input.Stats.Session.Battles == 0 {
		return stats.StatsBlock{}, fmt.Errorf("session stats 	have 0 battles")
	}

	var block stats.StatsBlock
	block.Tags = append(block.Tags, input.Options.Block.GenerationTag+"Block")
	block.Rows = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), 1, (input.Stats.AllTime.DamageDealt), 1)
	return block, nil
}
