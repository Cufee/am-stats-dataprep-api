package legacy

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func accuracy(allTime, label bool) *logic.Layout {
	var layout logic.Layout
	layout.Style = shared.AlignVertical.Merge(shared.JustifyCenter)
	// Session
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(shared.Gap10).Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
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
					Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
					Format:     "%v%%",
					Parse:      shared.RoundFloat,
				},
			},
			{ // Invisible icon to center things
				AddCondition: logic.SessionOfOverZero,
				Type:         logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { return smallIconSize },
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
						Expression: fmt.Sprintf("(%v / %v) * 100", logic.AllTimeValue, logic.AllTimeOf),
						Format:     "%v%%",
						Parse:      shared.RoundFloat,
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
						String:   "localized_shot_accuracy",
					},
				},
			},
		},
		)
	}

	return &layout
}
