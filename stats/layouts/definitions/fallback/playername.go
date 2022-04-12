package fallback

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

func playerName() *logic.Layout {
	var layout logic.Layout
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				Type: logic.ItemTypeText,
				Data: logic.Text{},
			},
		},
	})
	return &layout
}

func playerTag() *logic.Layout {
	var layout logic.Layout
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				Type: logic.ItemTypeText,
				Data: logic.Text{},
			},
		},
	})
	return &layout
}
