package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"

	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func BattlesBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	if input.Stats.AllTime.Battles == 0 {
		input.Options.WithAllTime = false
		input.Options.Block.HasIcon = false
	}
	fixTag := fmt.Sprintf("fixIcon-%v", input.Options.Block.HasIcon && input.Options.Block.HasInvisibleIcon)
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.Battles), 1, (input.Stats.AllTime.Battles), 1, fixTag)
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, input.Options.Block.GenerationTag+"Block"))
	b.ContentType = block.ContentTypeBlocks
	b.Tags = append(b.Tags, fixTag)
	return b, nil
}

func WinrateBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	if input.Stats.AllTime.Battles == 0 {
		input.Options.WithAllTime = false
		input.Options.Block.HasIcon = false
	}
	fixTag := fmt.Sprintf("fixIcon-%v", input.Options.Block.HasIcon && input.Options.Block.HasInvisibleIcon)
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, true, (input.Stats.Session.Wins), (input.Stats.Session.Battles), (input.Stats.AllTime.Wins), (input.Stats.AllTime.Battles), fixTag)
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, input.Options.Block.GenerationTag+"Block"))
	b.ContentType = block.ContentTypeBlocks
	b.Tags = append(b.Tags, fixTag)
	return b, nil
}

func WinrateWithBattlesBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	if input.Stats.AllTime.Battles == 0 {
		input.Options.WithAllTime = false
		input.Options.Block.HasIcon = false
	}
	var fmtString utils.FmtStr
	fmtString.Session = "%v" + fmt.Sprintf(" (%v)", (input.Stats.Session.Battles))
	fmtString.AllTime = "%v" + fmt.Sprintf(" (%v)", (input.Stats.AllTime.Battles))
	fixTag := fmt.Sprintf("fixIcon-%v", input.Options.Block.HasIcon && input.Options.Block.HasInvisibleIcon)
	b.Content = utils.PrepContentRows(input, fmtString, true, (input.Stats.Session.Wins), (input.Stats.Session.Battles), (input.Stats.AllTime.Wins), (input.Stats.AllTime.Battles), fixTag)
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, input.Options.Block.GenerationTag+"Block"))
	b.ContentType = block.ContentTypeBlocks
	b.Tags = append(b.Tags, fixTag)
	return b, nil
}
