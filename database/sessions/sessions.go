package sessions

import (
	"byvko.dev/repo/am-stats-dataprep-api/database/sessions/types"
	"github.com/byvko-dev/am-core/firebase/firestore/driver"
)

const sessionsCollection = "userSessions"

func FindSessionDetails(session string) (types.UserSession, error) {
	var sessionDetails types.UserSession

	client, err := driver.NewDriver()
	if err != nil {
		return sessionDetails, err
	}

	err = client.GetDocumentByID(sessionsCollection, session, &sessionDetails)
	if err != nil {
		return sessionDetails, err
	}

	return sessionDetails, nil
}
