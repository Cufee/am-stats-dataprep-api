package stats

import "github.com/byvko-dev/am-types/wargaming/v1/statistics"

// VehicleStats - Player Vehicle stats struct, used to return final data
type VehicleStats struct {
	statistics.StatsFrame `json:"all" bson:"all"`
	LastBattleTime        int    `json:"last_battle_time" bson:"last_battle_time"`
	MarkOfMastery         int    `json:"mark_of_mastery" bson:"mark_of_mastery"`
	TankID                int    `json:"tank_id" bson:"tank_id"`
	TankTier              int    `json:"tank_tier" bson:"tank_tier"`
	TankName              string `json:"tank_name" bson:"tank_name"`
	TankWN8               int    `json:"tank_wn8,omitempty" bson:"tank_wn8,omitempty"`
	TankRawWN8            int    `json:"tank_raw_wn8,omitempty" bson:"tank_raw_wn8,omitempty"`
}

// Diff - Calculate the difference in two VehicleStats structs
func Diff(oldStats VehicleStats, newStats VehicleStats) (diffStats VehicleStats) {
	// Set stats fields
	oldStats.StatsFrame.Substract(&newStats.StatsFrame)
	diffStats.StatsFrame = oldStats.StatsFrame
	// Set other fields
	diffStats.LastBattleTime = newStats.LastBattleTime
	diffStats.TankID = newStats.TankID
	diffStats.TankWN8 = 0
	diffStats.TankRawWN8 = 0
	return diffStats
}
