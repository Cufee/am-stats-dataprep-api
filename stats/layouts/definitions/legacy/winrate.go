package legacy

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func winrate(allTime, label bool) *logic.Layout {
	var layout logic.Layout
	layout.Style = shared.AlignVertical
	// Session
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(shared.Gap25).Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				Style: smallIconSize,
				Type:  logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { style, _ := percentageIconStyleAndName(values); return style },
					GetName:  func(values logic.Values) string { _, name := percentageIconStyleAndName(values); return name },
				},
			},
			{
				Type: logic.ItemTypeTemplate,
				Data: logic.Template{
					Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
					Format:     "%v%%",
					Parse:      shared.RoundFloat,
				},
			},
			{ // Invisible icon to center things
				Style: smallIconSize,
				Type:  logic.ItemTypeIcon,
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
						String:   "localized_winrate",
					},
				},
			},
		})
	}
	return &layout
}

func winrateWithBattles(allTime, label bool) *logic.Layout {
	var layout logic.Layout
	layout.Style = shared.AlignVertical.Merge(shared.JustifyCenter)
	// Session
	layout.Rows = append(layout.Rows, logic.LayoutRow{
		Style: textLarge.Merge(shared.Gap25).Merge(textLargeColor),
		Items: []logic.LayoutItem{
			{
				Style: smallIconSize,
				Type:  logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { style, _ := percentageIconStyleAndName(values); return style },
					GetName:  func(values logic.Values) string { _, name := percentageIconStyleAndName(values); return name },
				},
			},
			{
				Type: logic.ItemTypeTemplate,
				Data: logic.Template{
					Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
					Format:     "%v%%",
					Parse:      shared.RoundFloat,
				},
			},
			{
				Style: TextMediumColor.Merge(TextMedium),
				Type:  logic.ItemTypeTemplate,
				Data: logic.Template{
					Expression: fmt.Sprintf("%v", logic.SessionOf),
					Format:     " (%v)",
				},
			},
			{ // Invisible icon to center things
				Style: smallIconSize,
				Type:  logic.ItemTypeIcon,
				Data: logic.Icon{
					GetStyle: func(values logic.Values) style.Style { return smallIconSize },
					GetName:  func(values logic.Values) string { _, name := percentageIconStyleAndName(values); return name },
				},
			},
		},
	})
	// All Time
	if allTime {
		layout.Rows = append(layout.Rows, logic.LayoutRow{
			Style: textLarge.Merge(textLargeColor),
			Items: []logic.LayoutItem{
				{
					Type: logic.ItemTypeTemplate,
					Data: logic.Template{
						Expression: fmt.Sprintf("(%v / %v) * 100", logic.AllTimeValue, logic.AllTimeOf),
						Format:     "%v%%",
						Parse:      shared.RoundFloat,
					},
				},
				{
					Type: logic.ItemTypeTemplate,
					Data: logic.Template{
						Expression: fmt.Sprintf("%v", logic.AllTimeOf),
						Format:     " (%v)",
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
						String:   "localized_winrate_with_battles",
					},
				},
			},
		})
	}
	return &layout
}
