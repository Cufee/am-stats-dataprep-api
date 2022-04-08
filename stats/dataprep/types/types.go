package types

import (
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type DataprepInput struct {
	Localizer *i18n.Localizer
	Stats     struct {
		Session statistics.StatsFrame
		AllTime statistics.StatsFrame
	}
	Options struct {
		WithLabel         bool
		WithIcons         bool
		WithInvisibleIcon bool
		WithAllTime       bool
		Block             settings.BlockOptions
		Style             string
	}
}
