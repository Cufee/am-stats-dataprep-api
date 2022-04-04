package presets

import "github.com/byvko-dev/am-types/dataprep/v1/settings"

func GetPresetByName(name string) settings.Options {
	switch name {
	case "minimal":
		return MinimalOptions
	default:
		return DefaultOptions
	}
}
