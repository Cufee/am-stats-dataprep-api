package legacy

import (
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var cardStyle = style.Style{
	PaddingLeft:     1.5,
	PaddingRight:    1.5,
	PaddingTop:      0.5,
	PaddingBottom:   0.25,
	BackgroundColor: color.RGBA{30, 30, 30, 204},
	BorderRadius:    30,
}.Merge(shared.DefaultFont).Merge(shared.AlignVertical).Merge(shared.GrowX)

var contentBase = style.Style{
	// PaddingLeft:   1,
	PaddingBottom: 0.5,
	// PaddingRight:  -0.5,
}.Merge(shared.DefaultFont)

var vehicleSlimContentStyle = contentBase.Merge(style.Style{
	MinWidth:       550,
	JustifyContent: style.JustifyContentSpaceBetween,
	PaddingTop:     0.5,
})

var contentStyle = contentBase.Merge(shared.GrowX).Merge(shared.Gap25).Merge(style.Style{JustifyContent: style.JustifyContentSpaceBetween})
var overviewTextStyle = fallback.TextMedium.Merge(fallback.TextMediumColor)
var vehicleSlimCardStyle = cardStyle.Merge(style.Style{AlignItems: style.AlignItemsHorizontal, JustifyContent: style.JustifyContentSpaceBetween}).Merge(style.Style{PaddingLeft: 2, PaddingRight: 1.25, PaddingTop: 0.2, Gap: 1})
