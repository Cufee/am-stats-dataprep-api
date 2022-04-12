package legacy

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func wn8(allTime, label bool) *logic.Layout {
	var layout logic.Layout
	layout.Style = shared.AlignVertical
	// Session
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(shared.Gap15).Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				AddCondition: logic.SessionValueOverNegOne,
				Style:        baseIconSize,
				Type:         logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { style, _ := wn8IconStyleAndName(values); return style },
					GetName:  func(values logic.Values) string { _, name := wn8IconStyleAndName(values); return name },
				},
			},
			{
				AddCondition: logic.SessionValueOverNegOne,
				Type:         logic.ItemTypeTemplate,
				Data: logic.Template{
					Expression: fmt.Sprintf("%v", logic.SessionValue),
					Format:     "%v",
				},
			},
			{ // Invisible icon to center things
				AddCondition: logic.SessionValueOverNegOne,
				Style:        baseIconSize,
				Type:         logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { return baseIconSize },
					GetName:  func(values logic.Values) string { _, name := wn8IconStyleAndName(values); return name },
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
					AddCondition: logic.AllTimeValueOverNegOne,
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
					AddCondition: func(v logic.Values) bool { return logic.AllTimeValueOverNegOne(v) || logic.SessionValueOverNegOne(v) },
					Style:        textSmall,
					Type:         logic.ItemTypeText,
					Data: logic.Text{
						Localize: true,
						String:   "localized_wn8_rating",
					},
				},
			},
		})
	}
	return &layout
}
