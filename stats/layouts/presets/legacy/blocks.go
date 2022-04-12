package legacy

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

var blocksDefault = []logic.Definition{
	fallback.BattlesDetailed,
	fallback.AvgDamageDetailed,
	fallback.WinrateDetailed,
	fallback.WN8Detailed,
}
