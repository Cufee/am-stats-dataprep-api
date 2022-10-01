package legacy

import (
	"strings"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/definitions/fallback"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
)

var VehiclesDetailed = logic.CardLayout{
	Title: logic.Text{
		Style: overviewTextStyle,
	},
	Blocks:       blocksDefault,
	Limit:        3,
	CardStyle:    cardStyle,
	ContentStyle: contentStyle,
}

var VehiclesSlim = logic.CardLayout{
	Title: logic.Text{
		Style:    overviewTextStyle,
		Localize: true,
		Printer:  trimVehicleName,
	},
	Blocks: []logic.Definition{
		fallback.AvgDamage,
		fallback.WinrateWithBattles,
		fallback.WN8,
	},
	Limit:        3,
	CardStyle:    vehicleSlimCardStyle,
	ContentStyle: vehicleSlimContentStyle,
}

func trimVehicleName(s string) string {
	nameSlice := strings.Split(s, " ")
	name := nameSlice[0]
	for i := 1; i < len(nameSlice); i++ {
		if len(name+" "+nameSlice[i]) > 21 {
			name += "..."
			break
		}
		name += " " + nameSlice[i]
	}
	return name
}
