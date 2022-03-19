package settings

import (
	"encoding/json"
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/database/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
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

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return settings.CreateNewSettings(userId, string(dataBytes))
}

func CreateNewSettingsWithID(id, userId string, data types.GenerationSettings) error {
	err := data.Validate()
	if err != nil {
		return err
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return settings.CreateNewSettingsWithID(id, userId, string(dataBytes))
}

func UpdateSettingsByID(userId, id string, data types.GenerationSettings) error {
	err := data.Validate()
	if err != nil {
		return err
	}

	ownerId, err := settings.GetSettingsOwnerId(id)
	if err != nil {
		return err
	}

	if ownerId != userId {
		return fmt.Errorf("user %v is not owner of settings %v", userId, id)
	}

	return settings.ReplaceSettingsByID(id, data)
}
