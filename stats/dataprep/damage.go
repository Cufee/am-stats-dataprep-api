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
	if input.Stats.AllTime.Battles == 0 {
		input.Options.WithAllTime = false
		input.Options.Block.HasIcon = false
	}
	b.ContentType = block.ContentTypeBlocks
	fixTag := fmt.Sprintf("fixIcon-%v", input.Options.Block.HasIcon && input.Options.Block.HasInvisibleIcon)
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, input.Options.Block.GenerationTag+"Block"))
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), (input.Stats.Session.Battles), (input.Stats.AllTime.DamageDealt), (input.Stats.AllTime.Battles), fixTag)
	b.Tags = append(b.Tags, fixTag)
	return b, nil
}

func DamageDoneBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats 	have 0 battles")
	}

	var b block.Block
	if input.Stats.AllTime.DamageDealt == 0 {
		input.Options.WithAllTime = false
		input.Options.Block.HasIcon = false
	}
	b.ContentType = block.ContentTypeBlocks
	fixTag := fmt.Sprintf("fixIcon-%v", input.Options.Block.HasIcon && input.Options.Block.HasInvisibleIcon)
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, input.Options.Block.GenerationTag+"Block"))
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), 1, (input.Stats.AllTime.DamageDealt), 1, fixTag)
	b.Tags = append(b.Tags, fixTag)
	return b, nil
}
