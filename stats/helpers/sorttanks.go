package helpers

import (
	"sort"

	"github.com/byvko-dev/am-core/stats/ratings/wn8/v1"
	"github.com/byvko-dev/am-types/stats/v3"
)

// SortTanks - Sorting of vehicles
func SortTanks(vehicles []stats.VehicleStats, sortKey string) []stats.VehicleStats {
	// Sort based on passed key
	switch sortKey {
	case "+battles":
		sort.Slice(vehicles, func(i, j int) bool {
			return vehicles[i].Stats.Battles < vehicles[j].Stats.Battles
		})
	case "-battles":
		sort.Slice(vehicles, func(i, j int) bool {
			return vehicles[i].Stats.Battles > vehicles[j].Stats.Battles
		})
	case "+winrate":
		sort.Slice(vehicles, func(i, j int) bool {
			return (float64(vehicles[i].Stats.Wins) / float64(vehicles[i].Stats.Battles)) < (float64(vehicles[j].Stats.Wins) / float64(vehicles[j].Stats.Battles))
		})
	case "-winrate":
		sort.Slice(vehicles, func(i, j int) bool {
			return (float64(vehicles[i].Stats.Wins) / float64(vehicles[i].Stats.Battles)) > (float64(vehicles[j].Stats.Wins) / float64(vehicles[j].Stats.Battles))
		})
	case "+wn8":
		sort.Slice(vehicles, func(i, j int) bool {
			return absInt(vehicles[i].Ratings[wn8.WN8]) < absInt(vehicles[j].Ratings[wn8.WN8])
		})
	case "-wn8":
		sort.Slice(vehicles, func(i, j int) bool {
			return absInt(vehicles[i].Ratings[wn8.WN8]) > absInt(vehicles[j].Ratings[wn8.WN8])
		})
	case "+last_battle":
		sort.Slice(vehicles, func(i, j int) bool {
			return absInt(vehicles[i].LastBattleTime) < absInt(vehicles[j].LastBattleTime)
		})
	case "-last_battle":
		sort.Slice(vehicles, func(i, j int) bool {
			return absInt(vehicles[i].LastBattleTime) > absInt(vehicles[j].LastBattleTime)
		})
	case "+damage":
		sort.Slice(vehicles, func(i, j int) bool {
			return int(float64(vehicles[i].Stats.DamageDealt)/float64(vehicles[i].Stats.Battles)) < int(float64(vehicles[j].Stats.DamageDealt)/float64(vehicles[j].Stats.Battles))
		})
	case "-damage":
		sort.Slice(vehicles, func(i, j int) bool {
			return int(float64(vehicles[i].Stats.DamageDealt)/float64(vehicles[i].Stats.Battles)) > int(float64(vehicles[j].Stats.DamageDealt)/float64(vehicles[j].Stats.Battles))
		})
	case "relevance":
		sort.Slice(vehicles, func(i, j int) bool {
			return (absInt(vehicles[i].Ratings[wn8.WN8]) * vehicles[i].LastBattleTime * vehicles[i].Stats.Battles) > (absInt(vehicles[j].Ratings[wn8.WN8]) * vehicles[j].LastBattleTime * vehicles[j].Stats.Battles)
		})
	default:
		sort.Slice(vehicles, func(i, j int) bool {
			return absInt(vehicles[i].LastBattleTime) > absInt(vehicles[j].LastBattleTime)
		})
	}
	return vehicles
}

// absInt - Absolute value of an integer
func absInt(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}
