package legacy

import (
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var cardStyle = style.Style{
	PaddingLeft:     1,
	PaddingRight:    1,
	PaddingTop:      0.5,
	PaddingBottom:   0.5,
	BackgroundColor: color.RGBA{30, 30, 30, 204},
	BorderRadius:    30,
}.Merge(shared.DefaultFont).Merge(shared.AlignVertical).Merge(shared.GrowX)
var contentStyle = style.Style{
	PaddingLeft:   0.75,
	PaddingBottom: 0.5,
	Gap:           2,
}.Merge(shared.DefaultFont).Merge(shared.GrowX)

var overviewTextStyle = fallback.TextMedium.Merge(fallback.TextMediumColor)
