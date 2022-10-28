package logic

import (
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type Definition struct {
	Name      string `json:"name"`
	ValueKind `json:"valueKind"`
	Fallback  interface{} `json:"fallback"`
}

type CardLayout struct {
	Title        Text         `json:"title"`
	Blocks       []Definition `json:"blocks"`
	CardStyle    style.Style  `json:"cardStyle"`
	ContentStyle style.Style  `json:"contentStyle"`
	Limit        int          `json:"limit"`
}

type LayoutOptions struct {
	LayoutName           string      `json:"name"`
	VehiclesSort         string      `json:"vehiclesSort"`
	WrapperStyle         style.Style `json:"wrapperStyle"`
	AccountStatus        *CardLayout `json:"accountStatus"`
	Notifications        *CardLayout `json:"notifications"`
	Challenges           *CardLayout `json:"challenges"`
	PlayerInfo           *CardLayout `json:"playerInfo"`
	RatingOverview       *CardLayout `json:"ratingOverview"`
	RandomOverview       *CardLayout `json:"randomOverview"`
	VehiclesFullOverview *CardLayout `json:"vehiclesFullOverview"`
	VehiclesSlimOverview *CardLayout `json:"vehiclesSlimOverview"`
}
