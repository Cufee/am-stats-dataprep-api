package fallback

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func accuracy() *logic.Layout {
	return &logic.Layout{
		Style: style.Style{
			PaddingRight: -0.5, // To fix the invisible icons
		},
		Rows: []logic.LayoutRow{
			{ // Session
				Items: []logic.LayoutItem{
					{
						Type: logic.ItemTypeIcon,
						Data: logic.Icon{
							GetStyle: func(values logic.Values) style.Style { style, _ := iconStyleAndName(values); return style },
							GetName:  func(values logic.Values) string { _, name := iconStyleAndName(values); return name },
						},
					},
					{
						Type: logic.ItemTypeTemplate,
						Data: logic.Template{
							Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
							Format:     "%v%%",
						},
					},
					{ // Invisible icon to center things
						Type: logic.ItemTypeIcon,
						Data: logic.Icon{
							GetStyle: func(values logic.Values) style.Style { return baseIconSize },
							GetName:  func(values logic.Values) string { _, name := iconStyleAndName(values); return name },
						},
					},
				},
			},
			{ // AllTime
				Items: []logic.LayoutItem{
					{
						Type: logic.ItemTypeIcon,
						Data: logic.Icon{
							GetStyle: func(values logic.Values) style.Style { style, _ := iconStyleAndName(values); return style },
							GetName:  func(values logic.Values) string { _, name := iconStyleAndName(values); return name },
						},
					},
					{
						Type: logic.ItemTypeTemplate,
						Data: logic.Template{
							Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
							Format:     "%v%%",
						},
					},
					{ // Invisible icon to center things
						Type: logic.ItemTypeIcon,
						Data: logic.Icon{
							GetStyle: func(values logic.Values) style.Style { return baseIconSize },
							GetName:  func(values logic.Values) string { _, name := iconStyleAndName(values); return name },
						},
					},
				},
			},
			{ // Label
				Items: []logic.LayoutItem{
					{
						Style: textSmall,
						Type:  logic.ItemTypeText,
						Data: logic.Text{
							Localize: true,
							String:   "localized_shot_accuracy",
						},
					},
				},
			},
		},
	}
}
