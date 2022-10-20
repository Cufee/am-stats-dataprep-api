package settings

import (
	"time"

	"byvko.dev/repo/am-stats-dataprep-api/database"
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSettingsByID(id string) (*settings.GenerationSettings, error) {
	var data settings.GenerationSettings
	return &data, database.GetSettingsByID(id, &data)
}

func CreateNewSettings(userId string, data settings.GenerationSettings) (string, error) {
	err := data.Validate()
	if err != nil {
		return "", err
	}

	data.OwnerId = userId
	data.LastUsed = time.Now()
	dataBytes, err := bson.Marshal(data)
	if err != nil {
		return "", err
	}

	return database.CreateNewSettings(dataBytes)
}

func UpdateSettingsByID(id string, data settings.GenerationSettings) error {
	err := data.Validate()
	if err != nil {
		return err
	}
	// Encode with bson
	dataBytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}

	// Convert to map
	var dataMap map[string]interface{}
	err = bson.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return err
	}

	return database.UpdateSettingsByID(id, dataMap)
}
