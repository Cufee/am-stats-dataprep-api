package statsapi

import (
	"encoding/json"
	"fmt"

	"github.com/byvko-dev/am-core/helpers/env"
	"github.com/byvko-dev/am-core/helpers/requests"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/api/stats/v1"
)

var apiUrl = env.MustGetString("STATS_API_URL")

func GetStatsByPlayerID(playerID int, locale string, days int) (*stats.ResponsePayload, error) {
	if playerID == 0 {
		return nil, fmt.Errorf("playerID and realm are required")
	}

	var request stats.RequestPayload
	request.AccountID = playerID
	request.Locale = locale
	request.Days = days
	return GetStatsFromRequest(request)
}

func GetStatsFromRequest(request stats.RequestPayload) (*stats.ResponsePayload, error) {
	if request.AccountID == 0 {
		return nil, parseStatsError(fmt.Errorf("playerID and realm are required"))
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, parseStatsError(logs.Wrap(err, "failed to marshal request"))
	}

	var response stats.ResponsePayload
	_, err = requests.Send(fmt.Sprintf("%v/session/player", apiUrl), "POST", nil, payload, &response)
	if err != nil {
		return nil, parseStatsError(logs.Wrap(err, "Failed to get stats by player ID"))
	}
	return &response, nil
}
