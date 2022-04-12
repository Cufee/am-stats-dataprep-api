package presets

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets/legacy"
)

func LoadOPresetByName(name string) *logic.LayoutOptions {
	switch name {
	case "legacy":
		return &legacy.Preset
	default:
		return &fallback.Preset
	}
}

func Init() {
	fallback.Init()
	legacy.Init()
}
