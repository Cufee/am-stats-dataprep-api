package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func AvarageDamageBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, b.Tags...))
	b.ContentType = block.ContentTypeBlocks
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), (input.Stats.Session.Battles), (input.Stats.AllTime.DamageDealt), (input.Stats.AllTime.Battles))
	return b, nil
}

func DamageDoneBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats 	have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, b.Tags...))
	b.ContentType = block.ContentTypeBlocks
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), 1, (input.Stats.AllTime.DamageDealt), 1)
	return b, nil
}
