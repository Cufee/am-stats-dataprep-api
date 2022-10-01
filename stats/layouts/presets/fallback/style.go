package fallback

import (
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var cardStyle = style.Style{
	PaddingLeft:     0.5,
	PaddingRight:    0.5,
	PaddingTop:      0.5,
	PaddingBottom:   0.5,
	BackgroundColor: color.RGBA{30, 30, 30, 204},
	BorderRadius:    25,
}.Merge(shared.DefaultFont).Merge(shared.AlignVertical).Merge(shared.GrowX)
var contentStyle = style.Style{
	PaddingLeft:    1,
	PaddingRight:   0.5,
	PaddingBottom:  0.5,
	JustifyContent: style.JustifyContentSpaceBetween,
}.Merge(shared.DefaultFont).Merge(shared.Gap50).Merge(shared.GrowX)

var overviewTextStyle = fallback.TextMedium.Merge(fallback.TextMediumColor)
