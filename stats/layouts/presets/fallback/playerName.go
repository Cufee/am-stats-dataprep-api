package fallback

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

var PlayerName = logic.CardLayout{
	Blocks: []logic.Definition{
		fallback.PlayerName,
		fallback.PlayerClanTag,
	},
	CardStyle:    cardStyle,
	ContentStyle: shared.DefaultFont.Merge(shared.Gap50),
}
