package cache

// import (
// 	"time"

// 	"byvko.dev/repo/am-stats-dataprep-api/database/stats"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type statsCache struct {
// 	ID           primitive.ObjectID  `json:"id" firestore:"id" bson:"_id,omitempty"`
// 	Data         types.StatsResponse `json:"data" firestore:"data" bson:"data"`
// 	CreationTime time.Time           `json:"creationTime" firestore:"creationTime" bson:"creationTime"`
// }

// func GetStatsCacheByID(id string) (*types.StatsResponse, error) {
// 	var cache statsCache
// 	err := stats.GetStatsCacheByID(id, &cache)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &cache.Data, nil
// }

// func CreateStatsCache(data types.StatsResponse) (string, error) {
// 	var cache statsCache
// 	cache.Data = data
// 	cache.CreationTime = time.Now()
// 	return stats.CreateNewStatsCache(cache)
// }
