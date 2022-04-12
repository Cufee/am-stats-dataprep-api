package legacy

import "byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"

var (
	WN8 = logic.Definition{
		Name:      "wn8",
		ValueKind: logic.WN8OverOne,
	}
	WN8Detailed = logic.Definition{
		Name:      "wn8Detailed",
		ValueKind: logic.WN8OverOne,
	}
)
