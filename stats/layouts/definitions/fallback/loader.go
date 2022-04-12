package fallback

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

func Loader(name logic.Definition) *logic.Layout {
	switch name {
	case PlayerName:
		return playerName()
	case PlayerClanTag:
		return playerTag()

	case Battles:
		return battles(false, true)
	case BattlesDetailed:
		return battles(true, true)

	case Accuracy:
		return accuracy(false, true)
	case AccuracyDetailed:
		return accuracy(true, true)

	case Winrate:
		return winrate(false, true)
	case WinrateDetailed:
		return winrate(true, true)

	case WinrateWithBattles:
		return winrateWithBattles(false, true)
	case WinrateWithBattlesDetailed:
		return winrateWithBattles(true, true)

	case AvgDamage:
		return avgDamage(false, true)
	case AvgDamageDetailed:
		return avgDamage(true, true)

	case WN8:
		return wn8(false, true)
	case WN8Detailed:
		return wn8(true, true)
	default:
		return nil
	}
}
