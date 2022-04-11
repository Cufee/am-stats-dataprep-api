package fallback

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func iconStyleAndName(values logic.Values) (style.Style, string) {
	var iconStyle style.Style = baseIconSize
	var name string = shared.IconsLines[shared.IconDirectionHorizontal]

	// Make icon invisible if there is no data
	if val, ok := values[logic.SessionOf].(float64); !ok || val <= 0 {
		return iconStyle, name
	}

	iconStyle.BackgroundColor = shared.ColorNeutral // Start with neutral as baseline
	result, err := logic.EvaluateExpression(fmt.Sprintf("%v > %v", logic.AllTimeValue, logic.SessionValue), values)
	if err != nil {
		return shared.InvalidStyle, name
	}
	if result == "true" {
		iconStyle.BackgroundColor = shared.ColorGreen
		name = shared.IconsArrows[shared.IconDirectionUpSmall]
	} else {
		iconStyle.BackgroundColor = shared.ColorRed
		name = shared.IconsArrows[shared.IconDirectionDownSmall]
	}

	return iconStyle, name
}
