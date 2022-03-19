package statsapi

import (
	"github.com/byvko-dev/am-core/helpers"
)

var StatsApiUrl string
var StatsApiKey string

func init() {
	StatsApiKey, _ = helpers.MustGetEnv("STATS_API_KEY")[0].(string)
	StatsApiUrl, _ = helpers.MustGetEnv("STATS_API_URL")[0].(string)
}
