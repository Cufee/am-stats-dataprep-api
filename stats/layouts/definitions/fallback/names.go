package fallback

import "byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"

var (
	PlayerName = logic.Definition{
		Name:      "playerName",
		ValueKind: logic.PlayerName,
	}
	PlayerClanTag = logic.Definition{
		Name:      "playerClanTag",
		ValueKind: logic.PlayerClanTag,
	}

	BattlesDetailed = logic.Definition{
		Name:      "battlesDetailed",
		ValueKind: logic.BattlesOverOne,
	}
	Battles = logic.Definition{
		Name:      "battles",
		ValueKind: logic.BattlesOverOne,
	}

	AvgDamage = logic.Definition{
		Name:      "avgDamage",
		ValueKind: logic.DamageOverBattles,
	}
	AvgDamageDetailed = logic.Definition{
		Name:      "avgDamageDetailed",
		ValueKind: logic.DamageOverBattles,
	}

	Winrate = logic.Definition{
		Name:      "winrate",
		ValueKind: logic.WinsOverBattles,
	}
	WinrateDetailed = logic.Definition{
		Name:      "winrateDetailed",
		ValueKind: logic.WinsOverBattles,
	}

	WinrateWithBattles = logic.Definition{
		Name:      "winrateWithBattles",
		ValueKind: logic.WinsOverBattles,
	}

	WinrateWithBattlesDetailed = logic.Definition{
		Name:      "winrateWithBattlesDetailed",
		ValueKind: logic.WinsOverBattles,
	}

	Accuracy = logic.Definition{
		Name:      "accuracy",
		ValueKind: logic.HitsOverShots,
	}
	AccuracyDetailed = logic.Definition{
		Name:      "accuracyDetailed",
		ValueKind: logic.HitsOverShots,
	}

	WN8 = logic.Definition{
		Name:      "wn8",
		ValueKind: logic.WN8OverOne,
	}
	WN8Detailed = logic.Definition{
		Name:      "wn8Detailed",
		ValueKind: logic.WN8OverOne,
	}
)
