package fallback

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

var VehiclesDetailed = presets.Preset{
	Title: logic.Text{
		String:   "localized_random_overview_title",
		Localize: true,
	},
	Blocks: blocksDefault,
}

var VehiclesSlim = presets.Preset{
	Title: logic.Text{
		String:   "localized_random_overview_title",
		Localize: true,
	},
	Blocks: []shared.LayoutKind{
		shared.AvgDamage,
		shared.WinrateWithBattles,
		shared.Accuracy,
		shared.WN8,
	},
}
