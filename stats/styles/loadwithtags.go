package styles

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/legacy"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func LoadWithTags(styleName string, tags ...string) style.Style {
	switch styleName {
	case "minimal":
		return fallback.Load(tags...)

	case "legacy":
		return legacy.Load(tags...)

	default:
		return fallback.Load(tags...)
	}
}

func LoadBackground(styleName string, tags ...string) style.Style {
	var err error
	var bg style.Style
	switch styleName {
	case "minimal":
		bg, err = fallback.GetBackground(tags...)

	case "legacy":
		bg, err = legacy.GetBackground(tags...)

	default:
		bg, err = fallback.GetBackground(tags...)
	}

	if err != nil {
		image, _ := shared.LoadBackground("bg_default")
		return style.Style{
			BackgroundImage: image,
		}
	}
	return bg
}
