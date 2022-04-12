package shared

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var InvalidBlock = block.Block{
	ContentType: block.ContentTypeText,
	Content:     "invalid\nblock",
	Style: style.Style{
		Color:    color.RGBA{R: 255, G: 0, B: 0, A: 255},
		FontSize: 28,
	},
}
