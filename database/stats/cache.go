package stats

import (
	"github.com/byvko-dev/am-core/firebase/firestore/driver"
)

const settingsCollection = "stats/dataprep/cache"

func GetStatsCacheByID(id string, out interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}

	return driver.GetDocumentByID(settingsCollection, id, out)
}

func CreateNewStatsCache(data interface{}) (string, error) {
	driver, err := driver.NewDriver()
	if err != nil {
		return "", err
	}

	return driver.CreateDocumentInCollection(settingsCollection, data)
}
