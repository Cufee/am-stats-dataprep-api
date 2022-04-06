package styles

import (
	fallback "byvko.dev/repo/am-stats-dataprep-api/stats/styles/default"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func LoadWithTags(styleName string, tags ...string) style.Style {
	switch styleName {
	default:
		return fallback.Load(tags...)
	}
}
