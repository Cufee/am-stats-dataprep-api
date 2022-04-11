package layouts

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/legacy"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

func Load(name string, kind shared.LayoutKind) *logic.Layout {
	switch name {
	case "legacy":
		return legacy.Loader(kind)
	default:
		return fallback.Loader(kind)
	}
}
