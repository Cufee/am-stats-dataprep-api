package stats

import (
	"time"

	mongodb "github.com/byvko-dev/am-core/mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cacheTTL int32 = int32(time.Duration(5 * time.Minute).Seconds())
var collection = mongodb.NewCollection("stats-cache", []mongo.IndexModel{
	{
		Keys: bson.M{"creationTime": -1},
		Options: &options.IndexOptions{
			ExpireAfterSeconds: &cacheTTL,
		},
	},
})
