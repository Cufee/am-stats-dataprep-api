package realtime

import (
	"byvko.dev/repo/am-stats-dataprep-api/firebase/realtime/driver"
)

const settingsCollection = "stats/settings/"

func GetSettingsByID(id string, out interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}
	err = driver.GetDocumentByID(settingsCollection, id, out)
	if err != nil {
		return err
	}
	// defer func() {
	// 	update := make(map[string]interface{})
	// 	update["lastUsed"] = time.Now()
	// 	err = driver.UpdateDocumentByID(settingsCollection, id, update)
	// 	if err != nil {
	// 		logs.Error("Error updating lastUsed on %v: %v", id, err)
	// 	}
	// }()
	return nil
}

func CreateNewSettings(data interface{}) (string, error) {
	driver, err := driver.NewDriver()
	if err != nil {
		return "", err
	}
	return driver.CreateDocumentInCollection(settingsCollection, data)
}

func ReplaceSettingsByID(id string, data interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}
	return driver.ReplaceDocumentByID(settingsCollection, id, data)
}

func UpdateSettingsByID(id string, payload map[string]interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}
	return driver.UpdateDocumentByID(settingsCollection, id, payload)
}
