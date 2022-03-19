package types

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	statsTypes "github.com/byvko-dev/am-types/stats/v1"
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
