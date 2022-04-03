package settings

import (
	mongodb "github.com/byvko-dev/am-core/mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = mongodb.NewCollection("settings", []mongo.IndexModel{
	{Keys: bson.M{"ownerId": 1}}})
