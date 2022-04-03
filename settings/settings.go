package settings

import (
	"time"

	"byvko.dev/repo/am-stats-dataprep-api/database/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSettingsByID(id string) (*types.GenerationSettings, error) {
	var data types.GenerationSettings
	return &data, settings.GetSettingsByID(id, &data)
}

func CreateNewSettings(userId string, data types.GenerationSettings) (string, error) {
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

	return settings.CreateNewSettings(dataBytes)
}

// func CreateNewSettingsWithID(id, userId string, data types.GenerationSettings) error {
// 	err := data.Validate()
// 	if err != nil {
// 		return err
// 	}

// 	dataBytes, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	return settings.CreateNewSettingsWithID(id, userId, string(dataBytes))
// }

func UpdateSettingsByID(id string, data types.GenerationSettings) error {
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

	return settings.UpdateSettingsByID(id, dataMap)
}
