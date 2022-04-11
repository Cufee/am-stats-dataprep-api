package presets

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type Preset struct {
	Title  logic.Text          `json:"title"`
	Blocks []shared.LayoutKind `json:"blocks"`
	Style  style.Style         `json:"style"`
}

type Options struct {
	AccountStatus        bool `json:"accountStatus"`
	Notifications        bool `json:"notifications"`
	Challenges           bool `json:"challenges"`
	PlayerInfo           bool `json:"playerInfo"`
	RatingOverview       bool `json:"ratingOverview"`
	RandomOverview       bool `json:"randomOverview"`
	VehiclesFullOverview int  `json:"vehiclesFullOverview"`
	VehiclesSlimOverview int  `json:"vehiclesSlimOverview"`
}
