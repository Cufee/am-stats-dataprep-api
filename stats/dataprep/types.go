package dataprep

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type DataprepInput struct {
	Localizer *i18n.Localizer
	Stats     struct {
		Session *types.StatsFrame
		AllTime *types.StatsFrame
	}
	Options struct {
		WithLabel   bool
		WithAllTime bool
	}
}
