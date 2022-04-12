package layouts

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/legacy"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

func LoadDefinition(layoutName string, definition logic.Definition) *logic.Layout {
	switch layoutName {
	case "legacy":
		return legacy.Loader(definition)
	default:
		return fallback.Loader(definition)
	}
}
