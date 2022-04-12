package legacy

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

var OverviewRandom = logic.CardLayout{
	Title: logic.Text{
		Style:    overviewTextStyle,
		String:   "localized_random_overview_title",
		Localize: true,
	},
	Blocks:       blocksDefault,
	CardStyle:    cardStyle,
	ContentStyle: contentStyle,
}

var OverviewRating = logic.CardLayout{
	Title: logic.Text{
		Style:    overviewTextStyle,
		String:   "localized_rating_overview_title",
		Localize: true,
	},
	Blocks:       blocksDefault[:len(blocksDefault)-1],
	CardStyle:    cardStyle,
	ContentStyle: contentStyle,
}
