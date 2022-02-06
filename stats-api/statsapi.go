package statsapi

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"byvko.dev/repo/am-stats-dataprep-api/utils"
)

func GetStatsByPlayerID(playerID int, realm string, days int) (*types.PlayerRawStats, error) {
	var request types.StatsRequest
	request.PID = playerID
	request.Realm = realm
	request.Days = days

	var response types.PlayerRawStats

	var url = fmt.Sprintf("%v/stats", StatsApiUrl)
	err := utils.GetJSON(url, &response)

	return &response, err
}
