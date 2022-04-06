package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"

	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func ShotAccuracyBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	b.Style = styles.LoadWithTags(input.Options.Style, b.Tags...)
	b.ContentType = block.ContentTypeBlocks
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, b.Tags...))
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, true, (input.Stats.Session.Hits), (input.Stats.Session.Shots), (input.Stats.AllTime.Hits), (input.Stats.AllTime.Shots))
	return b, nil
}
