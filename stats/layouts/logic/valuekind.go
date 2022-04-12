package logic

type ValueKind string

const (
	HitsOverShots     ValueKind = "hitsOverShots"
	DamageOverBattles ValueKind = "damageOverBattles"
	WinsOverBattles   ValueKind = "winsOverBattles"
	BattlesOverOne    ValueKind = "battlesOverOne"
	WN8OverOne        ValueKind = "wn8OverOne"

	PlayerName    ValueKind = "playerName"
	PlayerClanTag ValueKind = "playerClanTag"
)
