package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"github.com/byvko-dev/am-types/dataprep/v1/block"
)

func AvarageDamageBlock(input types.DataprepInput) (block.Block, error) {
	if input.Stats.Session.Battles == 0 {
		return block.Block{}, fmt.Errorf("session stats have 0 battles")
	}

	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	b.Style.AlignItems = block.AlignItemsVertical
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
	b.Style.AlignItems = block.AlignItemsVertical
	b.ContentType = block.ContentTypeBlocks
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, (input.Stats.Session.DamageDealt), 1, (input.Stats.AllTime.DamageDealt), 1)
	return b, nil
}
