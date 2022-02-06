package types

// StatsFrame - Stats frame struct to unpack json and bson
type StatsFrame struct {
	Rating               float64 `json:"mm_rating" bson:"mm_rating"`
	Spotted              int     `json:"spotted" bson:"spotted"`
	Hits                 int     `json:"hits" bson:"hits"`
	Frags                int     `json:"frags" bson:"frags"`
	MaxXp                int     `json:"max_xp" bson:"max_xp"`
	Wins                 int     `json:"wins" bson:"wins"`
	Losses               int     `json:"losses" bson:"losses"`
	CapturePoints        int     `json:"capture_points" bson:"capture_points"`
	Battles              int     `json:"battles" bson:"battles"`
	DamageDealt          int     `json:"damage_dealt" bson:"damage_dealt"`
	DamageReceived       int     `json:"damage_received" bson:"damage_received"`
	MaxFrags             int     `json:"max_frags" bson:"max_frags"`
	Shots                int     `json:"shots" bson:"shots"`
	Xp                   int     `json:"xp" bson:"xp"`
	SurvivedBattles      int     `json:"survived_battles" bson:"survived_battles"`
	DroppedCapturePoints int     `json:"dropped_capture_points" bson:"dropped_capture_points"`
}

// FrameDiff - Calculate the difference in two StatsFrame structs
func FrameDiff(oldStats StatsFrame, newStats StatsFrame) (diffStats StatsFrame) {
	// Set stats fields
	diffStats.Spotted = newStats.Spotted - oldStats.Spotted
	diffStats.Hits = newStats.Hits - oldStats.Hits
	diffStats.Frags = newStats.Frags - oldStats.Frags
	diffStats.MaxXp = newStats.MaxXp - oldStats.MaxXp
	diffStats.Wins = newStats.Wins - oldStats.Wins
	diffStats.Losses = newStats.Losses - oldStats.Losses
	diffStats.CapturePoints = newStats.CapturePoints - oldStats.CapturePoints
	diffStats.Battles = newStats.Battles - oldStats.Battles
	diffStats.DamageDealt = newStats.DamageDealt - oldStats.DamageDealt
	diffStats.DamageReceived = newStats.DamageReceived - oldStats.DamageReceived
	diffStats.MaxFrags = newStats.MaxFrags - oldStats.MaxFrags
	diffStats.Shots = newStats.Shots - oldStats.Shots
	diffStats.Xp = newStats.Xp - oldStats.Xp
	diffStats.SurvivedBattles = newStats.SurvivedBattles - oldStats.SurvivedBattles
	diffStats.DroppedCapturePoints = newStats.DroppedCapturePoints - oldStats.DroppedCapturePoints

	return diffStats
}
