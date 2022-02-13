package statsapi

import (
	"encoding/json"
	"fmt"

	firebase "byvko.dev/repo/am-stats-dataprep-api/firebase/firestore"
	"byvko.dev/repo/am-stats-dataprep-api/helpers"
	"byvko.dev/repo/am-stats-dataprep-api/logs"
	"byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
)

const DefaultHeaderKeyIdentifier string = "x-api-key"

func GetStatsByPlayerID(playerID int, realm string, days int) (*types.PlayerRawStats, error) {
	headers := make(map[string]string)
	headers[DefaultHeaderKeyIdentifier] = StatsApiKey

	var request types.StatsRequest
	request.PID = playerID
	request.Realm = realm
	request.Days = days
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, logs.Wrap(err, "failed to marshal request")
	}

	var response types.PlayerRawStats
	var url = fmt.Sprintf("%v/stats", StatsApiUrl)
	_, err = helpers.HTTPRequest(url, "POST", headers, payload, &response)
	if err != nil {
		return nil, logs.Wrap(err, "Failed to get stats by player ID")
	}

	return &response, err
}

func GetMockStats() (*types.PlayerRawStats, error) {
	var response types.PlayerRawStats
	return &response, firebase.GetPreviewStats(&response)
}
