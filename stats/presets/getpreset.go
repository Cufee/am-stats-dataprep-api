package presets

import "github.com/byvko-dev/am-types/dataprep/settings/v1"

func GetPresetByName(name string) settings.Options {
	switch name {
	case "minimal":
		return MinimalOptions
	default:
		return DefaultOptions
	}
}
