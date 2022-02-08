package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/firebase"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
)

func GetSettingsByID(id string) (*types.GenerationSettings, error) {
	var settings types.GenerationSettings
	return &settings, firebase.GetSettingsByID(id, &settings)
}

func CreateNewSettings(data types.GenerationSettings) (string, error) {
	err := data.Validate()
	if err != nil {
		return "", err
	}
	return firebase.CreateNewSettings(data)
}

func UpdateSettingsByID(id string, data types.GenerationSettings) error {
	err := data.Validate()
	if err != nil {
		return err
	}
	return firebase.ReplaceSettingsByID(id, data)
}
