package shared

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var InvalidStyle = style.Style{
	BackgroundColor: color.RGBA{255, 0, 0, 255},
}
