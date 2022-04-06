package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	stats "byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func BlockFromStats(input types.DataprepInput) (block.Block, error) {
	switch input.Options.Block.GenerationTag {
	// Battles
	case stats.BlockBattles.GenerationTag:
		return BattlesBlock(input)

	case stats.BlockWinrate.GenerationTag:
		return WinrateBlock(input)

	case stats.BlockWinrateWithBattles.GenerationTag:
		return WinrateWithBattlesBlock(input)

		// Damage
	case stats.BlockDamageDone.GenerationTag:
		return BattlesBlock(input)

	case stats.BlockAverageDamage.GenerationTag:
		return AvarageDamageBlock(input)

	// Accuracy
	case stats.BlockShotAccuracy.GenerationTag:
		return ShotAccuracyBlock(input)

	default:
		return block.Block{}, fmt.Errorf("unknown block type: %s", input.Options.Block.GenerationTag)
	}
}
