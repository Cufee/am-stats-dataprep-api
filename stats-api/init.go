package statsapi

import "os"

var StatsApiUrl string
var StatsApiKey string

func init() {
	StatsApiKey = os.Getenv("STATS_API_KEY")
	StatsApiUrl = os.Getenv("STATS_API_URL")
	if StatsApiKey == "" || StatsApiUrl == "" {
		panic("STATS_API_KEY or STATS_API_URL not set")
	}
}
