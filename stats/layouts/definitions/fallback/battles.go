package fallback

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

func battles() *logic.Layout {
	return &logic.Layout{

		Rows: []logic.LayoutRow{
			{ // Session
				Items: []logic.LayoutItem{
					{
						Type: logic.ItemTypeTemplate,
						Data: logic.Template{
							Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
							Format:     "%v%%",
						},
					},
				},
			},
			{ // AllTime
				Items: []logic.LayoutItem{
					{
						Type: logic.ItemTypeTemplate,
						Data: logic.Template{
							Expression: fmt.Sprintf("(%v / %v) * 100", logic.SessionValue, logic.SessionOf),
							Format:     "%v%%",
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
							String:   "localized_battles",
						},
					},
				},
			},
		},
	}
}
