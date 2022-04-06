package styles

import (
	fallback "byvko.dev/repo/am-stats-dataprep-api/stats/styles/default"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func LoadWithTags(styleName string, tags ...string) style.Style {
	switch styleName {
	default:
		return fallback.Load(tags...)
	}
}

func LoadBackground(styleName string, tags ...string) style.Style {
	var err error
	var bg style.Style
	switch styleName {
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
