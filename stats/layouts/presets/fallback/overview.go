package fallback

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

var OverviewRandom = presets.Preset{
	Title: logic.Text{
		String:   "localized_random_overview_title",
		Localize: true,
	},
	Blocks: blocksDefault,
}

var OverviewRating = presets.Preset{
	Title: logic.Text{
		String:   "localized_rating_overview_title",
		Localize: true,
	},
	Blocks: []shared.LayoutKind{
		shared.Battles,
		shared.AvgDamage,
		shared.Winrate,
		shared.Accuracy,
	},
}
