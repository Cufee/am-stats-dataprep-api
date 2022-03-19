package cache

import (
	"byvko.dev/repo/am-stats-dataprep-api/database/stats"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func GetStatsCacheByID(id string) (*types.StatsResponse, error) {
	var data types.StatsResponse
	return &data, stats.GetStatsCacheByID(id, &data)
}

func CreateStatsCache(data types.StatsResponse) (string, error) {
	return stats.CreateNewStatsCache(data)
}
