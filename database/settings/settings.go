package settings

import (
	"encoding/json"
	"path"
	"time"

	"github.com/byvko-dev/am-core/firebase/firestore/driver"
	firebase "github.com/byvko-dev/am-core/firebase/types"
	"github.com/byvko-dev/am-core/logs"
)

const settingsCollection = "stats/dataprep/settings"

type databaseSettings struct {
	OwnerId string                `json:"ownerId" firestore:"ownerId"`
	Data    interface{}           `json:"data" firestore:"data"`
	Meta    firebase.DocumentMeta `json:"meta" firestore:"meta"`
}

func GetSettingsOwnerId(id string) (string, error) {
	driver, err := driver.NewDriver()
	if err != nil {
		return "", err
	}

	var owner string
	err = driver.GetDocumentByID(settingsCollection, path.Join(id, "ownerId"), &owner)
	if err != nil {
		return "", err
	}

	return owner, nil
}

func GetSettingsByID(id string, out interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}

	var document databaseSettings
	err = driver.GetDocumentByID(settingsCollection, id, &document)
	if err != nil {
		return err
	}

	defer func() {
		update := make(map[string]interface{})
		update["meta/lastUsed"] = time.Now()
		err = driver.UpdateDocumentByID(settingsCollection, id, update)
		if err != nil {
			logs.Error("Error updating lastUsed on %v: %v", id, err)
		}
	}()

	return json.Unmarshal([]byte(document.Data.(string)), out)
}

func CreateNewSettings(owner string, data interface{}) (string, error) {
	driver, err := driver.NewDriver()
	if err != nil {
		return "", err
	}

	var payload databaseSettings
	payload.Data = data
	payload.OwnerId = owner
	payload.Meta.CreationTime = time.Now()
	payload.Meta.LastUpdate = time.Now()
	payload.Meta.LastUsed = time.Now()

	return driver.CreateDocumentInCollection(settingsCollection, payload)
}

func CreateNewSettingsWithID(id, owner string, data interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}

	var payload databaseSettings
	payload.Data = data
	payload.OwnerId = owner
	payload.Meta.CreationTime = time.Now()
	payload.Meta.LastUpdate = time.Now()
	payload.Meta.LastUsed = time.Now()

	return driver.InsertDocument(settingsCollection, id, payload)
}

func ReplaceSettingsByID(id string, data interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}
	return driver.ReplaceDocumentByID(settingsCollection, path.Join(id, "data"), data)
}

func UpdateSettingsByID(id string, payload map[string]interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}
	return driver.UpdateDocumentByID(settingsCollection, path.Join(id, "data"), payload)
}
