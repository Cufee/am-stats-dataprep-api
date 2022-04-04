package statsapi

import (
	"github.com/byvko-dev/am-core/helpers/env"
)

var StatsApiUrl string
var StatsApiKey string

func init() {
	StatsApiKey, _ = env.MustGet("STATS_API_KEY")[0].(string)
	StatsApiUrl, _ = env.MustGet("STATS_API_URL")[0].(string)
}
