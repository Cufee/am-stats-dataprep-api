package statsapi

import (
	"encoding/json"
	"fmt"

	"github.com/byvko-dev/am-core/helpers/requests"
	"github.com/byvko-dev/am-core/logs"
	types "github.com/byvko-dev/am-types/stats/v1"
)

const DefaultHeaderKeyIdentifier string = "x-api-key"

func GetStatsByPlayerID(playerID int, realm string, days int) (*types.PlayerRawStats, error) {
	if playerID == 0 || realm == "" {
		return nil, fmt.Errorf("playerID and realm are required")
	}

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
	_, err = requests.Send(url, "POST", headers, payload, &response)
	if err != nil {
		return nil, logs.Wrap(err, "Failed to get stats by player ID")
	}

	return &response, err
}

func GetStatsFromRequest(request types.StatsRequest) (*types.PlayerRawStats, error) {
	if request.PID == 0 || request.Realm == "" {
		return nil, parseStatsError(fmt.Errorf("playerID and realm are required"))
	}

	headers := make(map[string]string)
	headers[DefaultHeaderKeyIdentifier] = StatsApiKey

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, parseStatsError(logs.Wrap(err, "failed to marshal request"))
	}

	var response types.PlayerRawStats
	var url = fmt.Sprintf("%v/stats", StatsApiUrl)
	_, err = requests.Send(url, "POST", headers, payload, &response)
	if err != nil {
		return nil, parseStatsError(logs.Wrap(err, "Failed to get stats by player ID"))
	}
	if response.Error != "" {
		return nil, parseStatsError(fmt.Errorf(response.Error))
	}

	return &response, nil
}
