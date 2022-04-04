package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"

	"github.com/byvko-dev/am-types/dataprep/v1/block"
)

func BattlesBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.Battles), 1, (input.Stats.AllTime.Battles), 1)
	b.Style.AlignItems = block.AlignItemsVertical
	b.ContentType = block.ContentTypeBlocks
	return b, nil
}

func WinrateBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, true, (input.Stats.Session.Wins), (input.Stats.Session.Battles), (input.Stats.AllTime.Wins), (input.Stats.AllTime.Battles))
	b.Style.AlignItems = block.AlignItemsVertical
	b.ContentType = block.ContentTypeBlocks
	return b, nil
}

func WinrateWithBattlesBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	var fmtString utils.FmtStr
	fmtString.Session = "%v" + fmt.Sprintf(" (%v)", (input.Stats.Session.Battles))
	fmtString.AllTime = "%v" + fmt.Sprintf(" (%v)", (input.Stats.AllTime.Battles))
	b.Content = utils.PrepContentRows(input, fmtString, true, (input.Stats.Session.Wins), (input.Stats.Session.Battles), (input.Stats.AllTime.Wins), (input.Stats.AllTime.Battles))
	b.Style.AlignItems = block.AlignItemsVertical
	b.ContentType = block.ContentTypeBlocks
	return b, nil
}
