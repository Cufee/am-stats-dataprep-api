package fallback

import (
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/helpers"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"
	"github.com/byvko-dev/am-core/helpers/slices"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var largeText = style.Style{
	Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
	FontSize: shared.DefaultFont.FontSize * 1.15,
}
var smallText = style.Style{
	Color:    color.RGBA{R: 140, G: 140, B: 140, A: 255},
	FontSize: shared.DefaultFont.FontSize * 0.8,
}
var mediumText = style.Style{
	Color: color.RGBA{R: 204, G: 204, B: 204, A: 255},
}
var card = style.Style{
	BackgroundColor: color.RGBA{
		R: 30,
		G: 30,
		B: 30,
		A: 204,
	},
	GrowX:         true,
	PaddingLeft:   0.75,
	PaddingRight:  0.75,
	PaddingTop:    0.5,
	PaddingBottom: 0.5,
	BorderRadius:  20,
}
var cardContent = shared.Gap50.Merge(shared.GrowX).Merge(shared.GrowY)
var wrapper = shared.Gap50.Merge(shared.Padding100)
var icon = style.Style{
	FontSize: shared.DefaultFont.FontSize * 0.5,
}
var iconOffsetPadding = style.Style{
	PaddingRight: -0.70,
}
var iconLarge = style.Style{
	FontSize: shared.DefaultFont.FontSize * 0.75,
}

func Load(tags ...string) style.Style {
	// Defaults / Fallbacks
	styleSheet := shared.DefaultFont
	styleSheet = shared.LoadStyles(styleSheet, tags...)

	// General layout
	if slices.Contains(tags, "card") > -1 {
		styleSheet = styleSheet.Merge(card).Merge(shared.Gap25)
	}
	if slices.Contains(tags, "wrapper") > -1 {
		styleSheet = styleSheet.Merge(wrapper)
	}
	if slices.Contains(tags, "content") > -1 {
		styleSheet = styleSheet.Merge(cardContent)
	}
	if helpers.SliceContains(tags, "statsContent", "vehicleOverview", "randomOverview", "ratingOverview", "playerNameContainer") {
		styleSheet = styleSheet.Merge(shared.GrowX).Merge(shared.Gap25)
	}

	// Gaps
	if slices.Contains(tags, "session") > -1 {
		styleSheet = styleSheet.Merge(shared.Gap25)
	}
	if helpers.SliceContains(tags, "statsContent", "label", "titleRow", "overviewTitle", "playerName") {
		styleSheet = styleSheet.Merge(shared.Gap25)
	}

	// Text
	if helpers.SliceContains(tags, "large", "session", "playerName") {
		styleSheet = styleSheet.Merge(largeText)
	}
	if helpers.SliceContains(tags, "allTime", "overviewTitle", "winrateWithBattlesBattles", "vehicleName") {
		styleSheet = styleSheet.Merge(mediumText)
	}
	if helpers.SliceContains(tags, "label", "vehicleTier") {
		styleSheet = styleSheet.Merge(smallText)
	}

	// Icon
	if slices.Contains(tags, "statsContent") > -1 && slices.Contains(tags, "fixIcon-true") > -1 {
		styleSheet = styleSheet.Merge(iconOffsetPadding)

	}
	if slices.Contains(tags, "icon") > -1 {
		if slices.Contains(tags, "wn8Rating") > -1 {
			styleSheet = styleSheet.Merge(iconLarge)
		} else {
			styleSheet = styleSheet.Merge(icon)
		}
	}

	// Left padding here is to fix font centering due to imperfect font metrics
	if slices.Contains(tags, "sessionValue") > -1 {
		styleSheet = styleSheet.Merge(style.Style{
			PaddingLeft: 0.075,
		})
	}

	// Background
	if slices.Contains(tags, "debug") > -1 {
		styleSheet = styleSheet.Merge(shared.DebugBackground)
	}
	return styleSheet
}

func GetBackground(tags ...string) (style.Style, error) {
	bg, err := shared.LoadBackground("bg_ukraine")
	if err != nil {
		return style.Style{}, err
	}

	return style.Style{
		BackgroundImage:     bg,
		BackgroundImageBlur: 15,
	}, nil
}
