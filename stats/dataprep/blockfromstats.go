package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func BlockFromStats(input DataprepInput, blockType string) (types.StatsBlock, error) {
	switch blockType {
	// Battles
	case types.BlockBattles:
		return BattlesBlock(input)

	case types.BlockWinrate:
		return WinrateBlock(input)

	case types.BlockWinrateWithBattles:
		return WinrateWithBattlesBlock(input)

		// Damage
	case types.BlockDamageDone:
		return BattlesBlock(input)

	case types.BlockAverageDamage:
		return AvarageDamageBlock(input)

	default:
		return types.StatsBlock{}, fmt.Errorf("unknown block type: %s", blockType)
	}
}
