package types

const (
	// Battles
	BlockBattles            = "battles"
	BlockBattlesWon         = "wins"
	BlockBattlesLost        = "losses"
	BlockBattlesSurvived    = "survived_battles"
	BlockWinrateWithBattles = "winrate_with_battles"
	BlockWinrate            = "winrate"

	// Damage
	BlockDamageDone     = "damage_dealt"
	BlockDamageReceived = "damage_received"
	BlockAverageDamage  = "average_damage"
	BlockDamageRatio    = "damage_ratio"

	// Frags / Shots
	BlockKills        = "frags"
	BlockMaxFrags     = "max_frags"
	BlockShotsHit     = "hits"
	BlockShotsFired   = "shots"
	BlockShotAccuracy = "shot_accuracy"

	// Assists
	BlockSpotted  = "spotted"
	BlockCaptures = "capture_points"
	BlockDrops    = "dropped_capture_points"

	// XP
	BlockXp    = "xp"
	BlockMaxXp = "max_xp"

	// Rating
	BlockAftermathRating = "aftermath_rating"
	BlockWargamingRating = "wargaming_rating"
	BlockWN8Rating       = "wn8_rating"

	// Challenges
	BlockChallengeTimePassed = "challenge_time_passed"
	BlockChallengeTimeLeft   = "challenge_time_left"

	BlockChallengeGoal     = "challenge_goal"
	BlockChallengeScore    = "challenge_wins"
	BlockChallengeProgress = "challenge_progress"

	BlockChallengeName   = "challenge_name"
	BlockChallengeSource = "challenge_source"

	// Challenge types
	ChallengeTypeAll           = "all"
	ChallengeTypeNone          = "none"
	ChallengeTypeCustom        = "custom"
	ChallengeTypeFromClan      = "from_clan"
	ChallengeTypeFromPromo     = "from_promo"
	ChallengeTypeFromPlayer    = "from_player"
	ChallengeTypeFromAftermath = "from_aftermath"
)
