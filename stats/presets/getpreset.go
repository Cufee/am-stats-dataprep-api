package presets

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets/legacy"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets/minimal"

	"github.com/byvko-dev/am-types/dataprep/settings/v1"
)

func GetPresetByName(name string) settings.Options {
	switch name {
	case "legacy":
		return legacy.Options
	case "minimal":
		return minimal.Options
	default:
		return fallback.Options
	}
}
