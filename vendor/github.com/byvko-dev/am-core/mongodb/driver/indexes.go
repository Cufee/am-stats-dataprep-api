package driver

import (
	"context"

	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-core/mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Client) validateIndexes(collection types.Collection) {
	go func(collection types.Collection) {
		newClient, err := NewClient()
		if err != nil {
			logs.Error("Failed to create new client for index validation: %v", err)
			return
		}
		db := newClient.Collection(collection.GetName())

		specs, err := db.Indexes().ListSpecifications(context.Background())
		if err != nil {
			logs.Debug("validateIndexes / Failed to list indexes: %v", err)
			return
		}

		if (len(specs) - 1) != len(collection.GetIndexes()) { // Specs always has _id index
			logs.Debug("validateIndexes / Refreshing all indexes for %v due to length %v != %v", collection.GetName(), len(specs), len(collection.GetIndexes()))
			// Refresh all indexes
			db.Indexes().CreateMany(context.Background(), collection.GetIndexes())
			db.Indexes().DropAll(context.Background())
			return
		}

		for _, index := range collection.GetIndexes() {
			found := false
			for _, spec := range specs {
				found = compareIndexKeys(spec, index)
				if found {
					break
				}
			}
			if !found {
				logs.Debug("validateIndexes / Refreshing all indexes for %v due to incorrect %v index", collection.GetName(), index)
				// Refresh all indexes
				db.Indexes().DropAll(context.Background())
				db.Indexes().CreateMany(context.Background(), collection.GetIndexes())
				return
			}
		}
		logs.Debug("validateIndexes / All indexes are valid for %v", collection.GetName())
	}(collection)
}

// True if matched, false otherwise
func compareIndexKeys(current *mongo.IndexSpecification, requested mongo.IndexModel) bool {
	currentKeys, err := bson.Marshal(current.KeysDocument)
	if err != nil {
		logs.Debug("compareIndexKeys / Failed to marshal keys: %v", err)
		return false
	}
	requestedKeys, err := bson.Marshal(requested.Keys)
	if err != nil {
		logs.Debug("compareIndexKeys / Failed to marshal keys: %v", err)
		return false
	}
	var keysMatch bool = string(currentKeys) == string(requestedKeys)
	if !keysMatch {
		return false
	}

	var expirationMatch = true
	if requested.Options != nil && requested.Options.ExpireAfterSeconds != nil {
		expirationMatch = current.ExpireAfterSeconds != nil && *current.ExpireAfterSeconds == *requested.Options.ExpireAfterSeconds
	} else {
		expirationMatch = current.ExpireAfterSeconds == nil
	}

	return keysMatch && expirationMatch
}
