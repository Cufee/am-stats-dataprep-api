package shared

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var CardBackground = style.Style{
	Color:         color.RGBA{R: 255, G: 255, B: 255, A: 255},
	PaddingLeft:   0.5,
	PaddingRight:  0.5,
	PaddingTop:    0.5,
	PaddingBottom: 0.5,
	BorderRadius:  30,
}
