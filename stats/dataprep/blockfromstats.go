package dataprep

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

func BlockFromStats(kind shared.LayoutKind, layoutName string, session, allTime statistics.StatsFrame) block.Block {
	switch kind {
	// // Battles
	// case shared.GenerationTagBattles:
	// 	return BattlesBlock(input)

	// case shared.GenerationTagWinrate:
	// 	return WinrateBlock(input)

	// case shared.GenerationTagWinrateWithBattles:
	// 	return WinrateWithBattlesBlock(input)

	// 	// Damage
	// case shared.GenerationTagDamageDone:
	// 	return DamageDoneBlock(input)

	// case shared.GenerationTagAverageDamage:
	// 	return AvarageDamageBlock(input)

	// Accuracy
	case shared.Accuracy:
		return ShotAccuracyBlock(layoutName, session, allTime)

	default:
		return shared.InvalidBlock
	}
}
