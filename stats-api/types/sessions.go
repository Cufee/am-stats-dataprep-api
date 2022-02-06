package types

import (
	"strconv"
	"time"
)

// Session - Will be switching to this format soon
type Session struct {
	Vehicles      []VehicleStats    `json:"vehicles" bson:"vehicles"`
	Achievements  AchievementsFrame `json:"achievements" bson:"achievements"`
	PlayerID      int               `json:"player_id" bson:"player_id"`
	Timestamp     time.Time         `json:"timestamp" bson:"timestamp"`
	LastBattle    time.Time         `json:"last_battle_time" bson:"last_battle_time"`
	BattlesAll    int               `json:"battles_random" bson:"battles_random"`
	StatsAll      StatsFrame        `json:"stats_random" bson:"stats_random"`
	BattlesRating int               `json:"battles_rating" bson:"battles_rating"`
	StatsRating   StatsFrame        `json:"stats_rating" bson:"stats_rating"`
	SessionRating int               `json:"session_wn8" bson:"session_wn8"`
	Convert       `json:"-" bson:"-"`
}

// RetroSession - Session using old data structure
type RetroSession struct {
	Vehicles      map[string]VehicleStats `json:"vehicles" bson:"vehicles"`
	Achievements  AchievementsFrame       `json:"achievements" bson:"achievements"`
	PlayerID      int                     `json:"player_id" bson:"player_id"`
	Timestamp     time.Time               `json:"timestamp" bson:"timestamp"`
	LastBattle    time.Time               `json:"last_battle_time" bson:"last_battle_time"`
	BattlesAll    int                     `json:"battles_random" bson:"battles_random"`
	StatsAll      StatsFrame              `json:"stats_random" bson:"stats_random"`
	BattlesRating int                     `json:"battles_rating" bson:"battles_rating"`
	StatsRating   StatsFrame              `json:"stats_rating" bson:"stats_rating"`
	SessionRating int                     `json:"session_wn8" bson:"session_wn8"`
	Convert       `json:"-" bson:"-"`
}

// Convert - Convert between Session and RetroSession
type Convert interface {
	ToSession() Session
	ToRetro() RetroSession
}

// ToSession - Covert RetroSession to Session Struct, Session is easier to work with in Go
func (s RetroSession) ToSession() (sessionNew Session) {
	sessionNew.Achievements = s.Achievements
	sessionNew.PlayerID = s.PlayerID
	sessionNew.Timestamp = s.Timestamp
	sessionNew.LastBattle = s.LastBattle
	sessionNew.BattlesAll = s.BattlesAll
	sessionNew.StatsAll = s.StatsAll
	sessionNew.BattlesRating = s.BattlesRating
	sessionNew.StatsRating = s.StatsRating
	sessionNew.SessionRating = s.SessionRating
	// Convert Vehicle Stats
	for _, v := range s.Vehicles {
		sessionNew.Vehicles = append(sessionNew.Vehicles, v)
	}
	return sessionNew
}

// ToRetro - Covert RetroSession to Session Struct, RetroSession is the format used by Aftermath rendering script.
func (s Session) ToRetro() (sessionNew RetroSession) {
	sessionNew.Achievements = s.Achievements
	sessionNew.PlayerID = s.PlayerID
	sessionNew.Timestamp = s.Timestamp
	sessionNew.LastBattle = s.LastBattle
	sessionNew.BattlesAll = s.BattlesAll
	sessionNew.StatsAll = s.StatsAll
	sessionNew.BattlesRating = s.BattlesRating
	sessionNew.StatsRating = s.StatsRating
	sessionNew.SessionRating = s.SessionRating
	// Convert Vehicle Stats
	vehicleMap := make(map[string]VehicleStats)
	for _, v := range s.Vehicles {
		vehicleMap[strconv.Itoa(v.TankID)] = v
	}
	sessionNew.Vehicles = vehicleMap
	return sessionNew
}
