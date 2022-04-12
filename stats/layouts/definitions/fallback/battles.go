package fallback

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

func battles(allTime, label bool) *logic.Layout {
	var layout logic.Layout
	layout.Style = shared.AlignVertical
	// Session
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				Type: logic.ItemTypeTemplate,
				Data: logic.Template{
					Expression: fmt.Sprintf("%v", logic.SessionValue),
					Format:     "%v",
				},
			},
		},
	})
	if allTime {
		// All Time
		layout.Rows = append(layout.Rows, logic.LayoutRow{
			Style: TextMedium.Merge(TextMediumColor),
			Items: []logic.LayoutItem{
				{
					AddCondition: logic.AllTimeValueOverZero,
					Type:         logic.ItemTypeTemplate,
					Data: logic.Template{
						Expression: fmt.Sprintf("%v", logic.AllTimeValue),
						Format:     "%v",
					},
				},
			},
		})
	}
	// Label
	if label {
		layout.Rows = append(layout.Rows, logic.LayoutRow{
			Style: textSmall.Merge(textSmallColor),
			Items: []logic.LayoutItem{
				{
					AddCondition: func(v logic.Values) bool {
						return (allTime && logic.AllTimeOfOverZero(v)) || logic.SessionOfOverZero(v)
					},
					Style: textSmall,
					Type:  logic.ItemTypeText,
					Data: logic.Text{
						Localize: true,
						String:   "localized_battles",
					},
				},
			},
		})
	}
	return &layout
}
