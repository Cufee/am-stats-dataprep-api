package firebase

import (
	"byvko.dev/repo/am-stats-dataprep-api/firebase/firestore/driver"
)

const statsCollection = "stats/sessions/"

func GetPreviewStats(out interface{}) error {
	driver, err := driver.NewDriver()
	if err != nil {
		return err
	}
	return driver.GetDocumentByID(statsCollection, "preview", out)
}
