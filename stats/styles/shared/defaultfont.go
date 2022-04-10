package shared

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var DefaultFont = style.Style{
	Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
	FontSize: 24,
}.Merge(FontRegular)

var FontRegular = style.Style{
	FontName: "roboto-regular",
}

var FontBold = style.Style{
	FontName: "roboto-bold",
}

var FontThin = style.Style{
	FontName: "roboto-thin",
}
