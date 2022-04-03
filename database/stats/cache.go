package stats

import (
	"errors"

	"github.com/byvko-dev/am-core/mongodb/driver"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStatsCacheByID(hex string, out interface{}) error {
	client, err := driver.NewClient()
	if err != nil {
		return err
	}
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}
	return client.GetDocumentWithFilter(collection, map[string]interface{}{"_id": id}, out)
}

func CreateNewStatsCache(data interface{}) (string, error) {
	client, err := driver.NewClient()
	if err != nil {
		return "", err
	}
	newId, err := client.InsertDocument(collection, data)
	if err != nil {
		return "", err
	}
	newIdString, ok := newId.(string)
	if ok {
		return newIdString, nil
	}
	newIdHex, ok := newId.(primitive.ObjectID)
	if ok {
		return newIdHex.Hex(), nil
	}
	return "", errors.New("unable to convert new id to string")
}
