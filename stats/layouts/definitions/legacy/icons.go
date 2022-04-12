package legacy

import (
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func wn8IconStyleAndName(values logic.Values) (style.Style, string) {
	var iconStyle style.Style = baseIconSize
	var name string = shared.IconCircle

	// Make icon invisible if there is no data
	if val, ok := values[logic.SessionValue].(float64); !ok || val <= 0 {
		return iconStyle, name
	}

	iconStyle.Color = getRatingColor(int(values[logic.SessionValue].(float64)))
	return iconStyle, name
}

// GetRatingColor - Rating color calculator
func getRatingColor(r int) color.RGBA {
	if r > 0 && r < 301 {
		return color.RGBA{255, 0, 0, 180}
	}
	if r > 300 && r < 451 {
		return color.RGBA{251, 83, 83, 180}
	}
	if r > 450 && r < 651 {
		return color.RGBA{255, 160, 49, 180}
	}
	if r > 650 && r < 901 {
		return color.RGBA{255, 244, 65, 180}
	}
	if r > 900 && r < 1201 {
		return color.RGBA{149, 245, 62, 180}
	}
	if r > 1200 && r < 1601 {
		return color.RGBA{103, 190, 51, 180}
	}
	if r > 1600 && r < 2001 {
		return color.RGBA{106, 236, 255, 180}
	}
	if r > 2000 && r < 2451 {
		return color.RGBA{46, 174, 193, 180}
	}
	if r > 2450 && r < 2901 {
		return color.RGBA{208, 108, 255, 180}
	}
	if r > 2900 {
		return color.RGBA{142, 65, 177, 180}
	}
	return color.RGBA{0, 0, 0, 0}
}
