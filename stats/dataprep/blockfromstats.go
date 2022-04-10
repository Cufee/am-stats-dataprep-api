package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets/shared"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func BlockFromStats(input types.DataprepInput) (block.Block, error) {
	switch input.Options.Block.GenerationTag {
	// Battles
	case shared.GenerationTagBattles:
		return BattlesBlock(input)

	case shared.GenerationTagWinrate:
		return WinrateBlock(input)

	case shared.GenerationTagWinrateWithBattles:
		return WinrateWithBattlesBlock(input)

		// Damage
	case shared.GenerationTagDamageDone:
		return DamageDoneBlock(input)

	case shared.GenerationTagAverageDamage:
		return AvarageDamageBlock(input)

	// Accuracy
	case shared.GenerationTagShotAccuracy:
		return ShotAccuracyBlock(input)

	default:
		return block.Block{}, fmt.Errorf("unknown block type: %s", input.Options.Block.GenerationTag)
	}
}
