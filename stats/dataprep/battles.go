package dataprep

import (
	"fmt"
	"log"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	statsTypes "byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func BattlesBlock(input types.DataprepInput) (statsTypes.StatsBlock, error) {
	if input.Stats.Session.Battles == 0 {
		return statsTypes.StatsBlock{}, fmt.Errorf("session stats have 0 battles")
	}

	var block statsTypes.StatsBlock
	block.Tags = append(block.Tags, input.Options.Block.GenerationTag+"Block")

	log.Print("Session ", (input.Stats.Session.Battles))
	log.Print("All ", (input.Stats.AllTime.Battles))

	block.Rows = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.Battles), 1, (input.Stats.AllTime.Battles), 1)
	return block, nil
}

func WinrateBlock(input types.DataprepInput) (statsTypes.StatsBlock, error) {
	if input.Stats.Session.Battles == 0 {
		return statsTypes.StatsBlock{}, fmt.Errorf("session stats have 0 battles")
	}

	var block statsTypes.StatsBlock
	block.Tags = append(block.Tags, input.Options.Block.GenerationTag+"Block")
	block.Rows = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, true, (input.Stats.Session.Wins), (input.Stats.Session.Battles), (input.Stats.AllTime.Wins), (input.Stats.AllTime.Battles))
	return block, nil
}

func WinrateWithBattlesBlock(input types.DataprepInput) (statsTypes.StatsBlock, error) {
	if input.Stats.Session.Battles == 0 {
		return statsTypes.StatsBlock{}, fmt.Errorf("session stats have 0 battles")
	}

	var block statsTypes.StatsBlock
	block.Tags = append(block.Tags, input.Options.Block.GenerationTag+"Block")
	var fmtString utils.FmtStr
	fmtString.Session = "%v" + fmt.Sprintf(" (%v)", (input.Stats.Session.Battles))
	fmtString.AllTime = "%v" + fmt.Sprintf(" (%v)", (input.Stats.AllTime.Battles))
	block.Rows = utils.PrepContentRows(input, fmtString, true, (input.Stats.Session.Wins), (input.Stats.Session.Battles), (input.Stats.AllTime.Wins), (input.Stats.AllTime.Battles))
	return block, nil
}
