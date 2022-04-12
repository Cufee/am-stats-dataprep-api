package legacy

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var baseFontSize = style.Style{
	FontSize: 28,
}
var baseIconSize = style.Style{
	FontSize: baseFontSize.FontSize * 0.75,
}
var textLarge = style.Style{
	FontSize: baseFontSize.FontSize * 1.25,
}
var textLargeColor = style.Style{
	Color: color.RGBA{255, 255, 255, 255},
}
var TextMedium = style.Style{
	FontSize: baseFontSize.FontSize * 1,
}
var TextMediumColor = style.Style{
	Color: color.RGBA{204, 204, 204, 255},
}
var textSmall = style.Style{
	FontSize: baseFontSize.FontSize * 0.75,
}
var textSmallColor = style.Style{
	Color: color.RGBA{100, 100, 100, 255},
}
