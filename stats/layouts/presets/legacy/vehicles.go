package legacy

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

var VehiclesDetailed = logic.CardLayout{
	Title: logic.Text{
		Style: overviewTextStyle,
	},
	Blocks:       blocksDefault,
	Limit:        3,
	CardStyle:    cardStyle,
	ContentStyle: contentStyle,
}

var VehiclesSlim = logic.CardLayout{
	Title: logic.Text{
		Style: overviewTextStyle,
	},
	Blocks: []logic.Definition{
		fallback.AvgDamage,
		fallback.WinrateWithBattles,
		fallback.WN8,
	},
	Limit:        3,
	CardStyle:    cardStyle,
	ContentStyle: contentStyle,
}
