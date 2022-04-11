package legacy

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

func Loader(kind shared.LayoutKind) *logic.Layout {
	switch kind {
	default:
		return fallback.Loader(kind)
	}
}
