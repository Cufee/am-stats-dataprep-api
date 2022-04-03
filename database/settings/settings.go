package settings

import (
	"errors"

	"github.com/byvko-dev/am-core/mongodb/driver"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSettingsByID(hex string, out interface{}) error {
	client, err := driver.NewClient()
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}

	filter := map[string]interface{}{"_id": id}
	return client.GetDocumentWithFilter(collection, filter, out)
}

func CreateNewSettings(data interface{}) (string, error) {
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

// func CreateNewSettingsWithID(id, data interface{}) error {
// 	client, err := driver.NewClient()
// 	if err != nil {
// 		return err
// 	}

// 	return driver.InsertDocument(settingsCollection, id, payload)
// }

// func ReplaceSettingsByID(id string, data interface{}) error {
// 	client, err := driver.NewClient()
// 	if err != nil {
// 		return err
// 	}

// 	client.UpdateDocumentWithFilter()

// 	return driver.ReplaceDocumentByID(settingsCollection, path.Join(id, "data"), data)
// }

func UpdateSettingsByID(hex string, payload map[string]interface{}) error {
	client, err := driver.NewClient()
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}
	filter := map[string]interface{}{"_id": id}
	return client.UpdateDocumentWithFilter(collection, filter, payload, false)
}
