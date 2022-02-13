package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	statsTypes "byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func ShotAccuracyBlock(input types.DataprepInput) (statsTypes.StatsBlock, error) {
	if input.Stats.Session.Battles == 0 {
		return statsTypes.StatsBlock{}, fmt.Errorf("session stats have 0 battles")
	}

	var block statsTypes.StatsBlock
	block.Tags = append(block.Tags, input.Options.Block.GenerationTag+"Block")
	block.Rows = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, true, (input.Stats.Session.Hits), (input.Stats.Session.Shots), (input.Stats.AllTime.Hits), (input.Stats.AllTime.Shots))
	return block, nil
}
