package legacy

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

func Loader(name logic.Definition) *logic.Layout {
	switch name {
	case WN8:
		return wn8(false, true)
	case WN8Detailed:
		return wn8(true, true)

	default:
		return fallback.Loader(name)
	}

}
