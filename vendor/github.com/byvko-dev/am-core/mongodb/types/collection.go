package types

import "go.mongodb.org/mongo-driver/mongo"

type collectionWithIndex struct {
	name    string             `bson:"name"`
	indexes []mongo.IndexModel `bson:"indexes"`
}

func (c *collectionWithIndex) GetName() string {
	return c.name
}
func (c *collectionWithIndex) GetIndexes() []mongo.IndexModel {
	return c.indexes
}

func NewCollection(name string, indexes []mongo.IndexModel) *collectionWithIndex {
	return &collectionWithIndex{
		name:    name,
		indexes: indexes,
	}
}

type Collection interface {
	GetName() string
	GetIndexes() []mongo.IndexModel
}
