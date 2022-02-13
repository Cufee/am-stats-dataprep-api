package types

import "byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/icons"

type BlockOptions struct {
	GenerationTag   string `json:"generationTag,omitempty"`
	LocalizationTag string `json:"localizationTag,omitempty"`

	IconDictOverwrite  map[string]string `json:"iconDictOverwrite,omitempty"`  // By default icons are arrows
	IconColorOverWrite string            `json:"iconColorOverwrite,omitempty"` // By default, icons are red,green,yellow based on session vs all time comparison
	HasIcon            bool              `json:"hasIcon,omitempty"`
}

func (b BlockOptions) WithIcon() BlockOptions {
	b.HasIcon = true
	return b
}

var (
	// Battles
	BlockBattles = BlockOptions{GenerationTag: "battles", LocalizationTag: "localized_battles"}

	// BlockBattlesWon         = "wins"
	// BlockBattlesLost        = "losses"
	// BlockBattlesSurvived    = "survived_battles"
	BlockWinrateWithBattles = BlockOptions{GenerationTag: "winrateWithBattles", LocalizationTag: "localized_winrate_with_battles"}
	BlockWinrate            = BlockOptions{GenerationTag: "winrate", LocalizationTag: "localized_winrate"}

	// Damage
	BlockDamageDone = BlockOptions{GenerationTag: "damageDone", LocalizationTag: "localized_damage_done"}
	// BlockDamageReceived = "damage_received"
	BlockAverageDamage = BlockOptions{GenerationTag: "averageDamage", LocalizationTag: "localized_average_damage"}
	// BlockDamageRatio    = "damage_ratio"

	// Frags / Shots
	// BlockKills        = "frags"
	// BlockMaxFrags     = "max_frags"
	// BlockShotsHit     = "hits"
	// BlockShotsFired   = "shots"
	BlockShotAccuracy = BlockOptions{GenerationTag: "shotAccuracy", LocalizationTag: "localized_shot_accuracy"}

	// Assists
	// BlockSpotted  = "spotted"
	// BlockCaptures = "capture_points"
	// BlockDrops    = "dropped_capture_points"

	// XP
	// BlockXp    = "xp"
	// BlockMaxXp = "max_xp"

	// Rating
	// BlockAftermathRating = "aftermath_rating"
	// BlockWargamingRating = "wargaming_rating"
	BlockWN8Rating = BlockOptions{GenerationTag: "wn8Rating", LocalizationTag: "localized_wn8_rating", IconDictOverwrite: icons.IconsCircle}

	// Challenges
	// BlockChallengeTimePassed = "challenge_time_passed"
	// BlockChallengeTimeLeft   = "challenge_time_left"

	// BlockChallengeGoal     = "challenge_goal"
	// BlockChallengeScore    = "challenge_wins"
	// BlockChallengeProgress = "challenge_progress"

	// BlockChallengeName   = "challenge_name"
	// BlockChallengeSource = "challenge_source"

	// Challenge types
	// ChallengeTypeAll           = "all"
	// ChallengeTypeNone          = "none"
	// ChallengeTypeCustom        = "custom"
	// ChallengeTypeFromClan      = "from_clan"
	// ChallengeTypeFromPromo     = "from_promo"
	// ChallengeTypeFromPlayer    = "from_player"
	// ChallengeTypeFromAftermath = "from_aftermath"
)
