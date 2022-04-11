package fallback

import "github.com/byvko-dev/am-types/dataprep/style/v1"

var baseFontSize = style.Style{
	FontSize: 28,
}
var baseIconSize = style.Style{
	FontSize: baseFontSize.FontSize * 0.5,
}
var largeSmall = style.Style{
	FontSize: baseFontSize.FontSize * 1.25,
}
var mediumSmall = style.Style{
	FontSize: baseFontSize.FontSize * 1,
}
var textSmall = style.Style{
	FontSize: baseFontSize.FontSize * 0.75,
}
