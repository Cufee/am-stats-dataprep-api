package fallback

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func avgDamage(allTime, label bool) *logic.Layout {
	var layout logic.Layout
	layout.Style = shared.AlignVertical
	// Session
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(shared.Gap10).Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				Style:        baseIconSize,
				AddCondition: logic.SessionOfOverZero,
				Type:         logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { style, _ := percentageIconStyleAndName(values); return style },
					GetName:  func(values logic.Values) string { _, name := percentageIconStyleAndName(values); return name },
				},
			},
			{
				AddCondition: logic.SessionOfOverZero,
				Type:         logic.ItemTypeTemplate,
				Data: logic.Template{
					Expression: fmt.Sprintf("(%v / %v)", logic.SessionValue, logic.SessionOf),
					Format:     "%v",
					Parse:      shared.FloatToInt,
				},
			},
			{ // Invisible icon to center things
				Style:        baseIconSize,
				AddCondition: logic.SessionOfOverZero,
				Type:         logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { return baseIconSize },
					GetName:  func(values logic.Values) string { _, name := percentageIconStyleAndName(values); return name },
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
					AddCondition: logic.AllTimeOfOverZero,
					Type:         logic.ItemTypeTemplate,
					Data: logic.Template{
						Expression: fmt.Sprintf("(%v / %v)", logic.AllTimeValue, logic.AllTimeOf),
						Format:     "%v",
						Parse:      shared.FloatToInt,
					},
				},
			},
		},
		)
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
						String:   "localized_average_damage",
					},
				},
			},
		},
		)
	}

	return &layout
}
