package fallback

import (
	"fmt"
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/helpers"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func wn8IconStyleAndName(values logic.Values) (style.Style, string) {
	var iconStyle style.Style = baseIconSize
	var name string = shared.IconsRating(3)[shared.IconDirectionUpLarge]

	// Make icon invisible if there is no data
	if val, ok := values[logic.SessionValue].(float64); !ok || val <= 0 {
		return iconStyle, name
	}

	clr, level := getRatingColorLevel(int(values[logic.SessionValue].(float64)))
	iconStyle.Color = clr
	name = shared.IconsRating(level)[shared.IconDirectionUpLarge]

	return iconStyle, name
}

func percentageIconStyleAndName(values logic.Values) (style.Style, string) {
	var iconStyle style.Style = smallIconSize
	var name string = shared.IconsLines[shared.IconDirectionHorizontal]

	// Make icon invisible if there is no data
	if val, ok := values[logic.SessionOf].(float64); !ok || val <= 0 {
		return iconStyle, name
	}

	iconStyle.Color = shared.ColorNeutral // Start with neutral as baseline
	result, err := logic.EvaluateExpression(fmt.Sprintf("(%v/%v) < (%v/%v)", logic.AllTimeValue, logic.AllTimeOf, logic.SessionValue, logic.SessionOf), values)
	if err != nil {
		return iconStyle, name
	}
	if result == "true" {
		// iconStyle.Color = shared.ColorGreen
		name = shared.IconsArrows[shared.IconDirectionUpSmall]
	} else {
		// iconStyle.Color = shared.ColorRed
		name = shared.IconsArrows[shared.IconDirectionDownSmall]
	}

	return iconStyle, name
}

// GetRatingColor - Rating color calculator
func getRatingColorLevel(r int) (color.RGBA, int) {
	// Red
	if r > 0 && r < 301 {
		return helpers.HexToColor("#fb5353"), 1
	}
	if r > 300 && r < 451 {
		return helpers.HexToColor("#fb5353"), 2
	}
	// Yellow
	if r > 450 && r < 651 {
		return helpers.HexToColor("#ffa02f"), 1
	}
	if r > 650 && r < 901 {
		return helpers.HexToColor("#ffa02f"), 2
	}
	// Green
	if r > 900 && r < 1201 {
		return helpers.HexToColor("#67be35"), 1
	}
	if r > 1200 && r < 1601 {
		return helpers.HexToColor("#67be35"), 2
	}
	// Teal
	if r > 1600 && r < 2001 {
		return helpers.HexToColor("#6ae6ff"), 1
	}
	if r > 2000 && r < 2451 {
		return helpers.HexToColor("#6ae6ff"), 2
	}
	// Purple / Pink
	if r > 2450 && r < 2901 {
		return helpers.HexToColor("#b757f7"), 1
	}
	if r > 2900 && r < 4501 {
		return helpers.HexToColor("#b757f7"), 2
	}
	// Legendary
	if r > 4500 && r < 6001 {
		return helpers.HexToColor("#f757b7"), 1
	}
	if r > 6000 {
		return helpers.HexToColor("#ffe84f"), 1
	}
	return color.RGBA{255, 255, 255, 0}, 1
}
