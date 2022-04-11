package fallback

import "byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"

var blocksDefault = []shared.LayoutKind{
	shared.Battles,
	shared.AvgDamage,
	shared.Winrate,
	shared.Accuracy,
	shared.WN8,
}
