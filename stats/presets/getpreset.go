package presets

import "byvko.dev/repo/am-stats-dataprep-api/stats/types"

func GetPresetByName(name string) types.Options {
	switch name {
	case "minimal":
		return MinimalOptions
	case "opus":
		return OpusOptions
	default:
		return DefaultOptions
	}
}
