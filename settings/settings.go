package settings

import (
	firebase "byvko.dev/repo/am-stats-dataprep-api/firebase/realtime"
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
	// payload := make(map[string]interface{})
	return firebase.ReplaceSettingsByID(id, data)
}
