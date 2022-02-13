package types

import (
	statsTypes "byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type DataprepInput struct {
	Localizer *i18n.Localizer
	Stats     struct {
		Session statsTypes.StatsFrame
		AllTime statsTypes.StatsFrame
	}
	Options struct {
		WithLabel   bool
		WithIcons   bool
		WithAllTime bool
		Block       types.BlockOptions
	}
}
