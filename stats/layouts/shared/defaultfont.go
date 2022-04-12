package shared

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var DefaultFont = style.Style{
	Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
	FontName: "roboto",
	FontSize: 28,
}
