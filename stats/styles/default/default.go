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
var statsContent = style.Style{
	PaddingLeft: 0.75,
}
var card = style.Style{
	BackgroundColor: color.RGBA{
		R: 30,
		G: 30,
		B: 30,
		A: 204,
	},
	GrowX:         true,
	PaddingLeft:   0.5,
	PaddingRight:  0.5,
	PaddingTop:    0.5,
	PaddingBottom: 0.5,
	BorderRadius:  20,
}
var cardContent = shared.Gap50.Merge(shared.GrowX).Merge(shared.GrowY)
var wrapper = shared.Gap50.Merge(shared.Padding100)
var icon = style.Style{
	FontSize: shared.DefaultFont.FontSize * 0.5,
}
var iconLarge = style.Style{
	FontSize: shared.DefaultFont.FontSize * 0.75,
}

func Load(tags ...string) style.Style {
	style := shared.DefaultFont
	if slices.Contains(tags, "icon") > -1 {
		if slices.Contains(tags, "wn8Rating") > -1 {
			style = style.Merge(iconLarge)
		} else {
			style = style.Merge(icon)
		}
	}
	if slices.Contains(tags, "large") > -1 {
		style = style.Merge(largeText)
	}
	if slices.Contains(tags, "gap25") > -1 {
		style = style.Merge(shared.Gap25)
	}
	if slices.Contains(tags, "gap50") > -1 {
		style = style.Merge(shared.Gap50)
	}
	if slices.Contains(tags, "growX") > -1 {
		style = style.Merge(shared.GrowX)
	}
	if slices.Contains(tags, "growY") > -1 {
		style = style.Merge(shared.GrowY)
	}
	if slices.Contains(tags, "card") > -1 {
		style = style.Merge(card).Merge(shared.Gap25)
	}
	if slices.Contains(tags, "wrapper") > -1 {
		style = style.Merge(wrapper)
	}
	if slices.Contains(tags, "content") > -1 {
		style = style.Merge(cardContent)
	}
	if slices.Contains(tags, "session") > -1 {
		style = style.Merge(largeText).Merge(shared.Gap25)
	}
	if slices.Contains(tags, "player_name") > -1 {
		style = style.Merge(shared.Gap25)
	}
	if slices.Contains(tags, "statsContent") > -1 {
		style = style.Merge(statsContent)
	}

	if helpers.SliceContains(tags, []string{"alltime", "overview_title", "winrate_with_battles_battles", "vehicleName"}) {
		style = style.Merge(mediumText)
	}
	if helpers.SliceContains(tags, []string{"label", "vehicleTier"}) {
		style = style.Merge(smallText)
	}
	if helpers.SliceContains(tags, []string{"label", "title_row", "overview_title"}) {
		style = style.Merge(shared.Gap25)
	}

	if slices.Contains(tags, "debug") > -1 {
		style = style.Merge(shared.DebugBackground)
	}
	return style
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
