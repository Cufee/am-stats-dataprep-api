package types

type Options struct {
	Locale string `json:"locale"`

	// Status icons
	AccountStatus StatusIconsOptions `json:"accountStatus,omitempty" firestore:"accountStatus"`

	// Notifications
	Notifications NotificationsOptions `json:"notifications,omitempty" firestore:"notifications"`

	// Challenges
	Challenges ChallengesOptions `json:"challenges,omitempty" firestore:"challenges"`

	// Player info and general stats
	Player         PlayerOptions   `json:"player,omitempty" firestore:"player"`
	RatingBattles  OverviewOptions `json:"ratingBattles,omitempty" firestore:"ratingBattles"`
	RegularBattles OverviewOptions `json:"regularBattles,omitempty" firestore:"regularBattles"`

	// Per vehicle stats
	VehiclesFull VehicleOptions `json:"vehiclesFull,omitempty" firestore:"vehiclesFull"`
	VehiclesSlim VehicleOptions `json:"vehiclesSlim,omitempty" firestore:"vehiclesSlim"`
}

type StatusIconsOptions struct {
	Include bool `json:"include,omitempty" firestore:"include"`
	Limit   int  `json:"limit,omitempty" firestore:"limit"`
}

type NotificationsOptions struct {
	Include bool     `json:"include,omitempty" firestore:"include"`
	Blocks  []string `json:"blocks,omitempty" 	firestore:"blocks"`
}

type ChallengesOptions struct {
	Include bool     `json:"include,omitempty" firestore:"include"`
	Limit   int      `json:"limit,omitempty" firestore:"limit"`
	Blocks  []string `json:"blocks,omitempty" firestore:"blocks"`
	Types   []string `json:"types,omitempty" firestore:"types"`
}

type PlayerOptions struct {
	Include     bool `json:"include,omitempty" firestore:"include"`
	WithName    bool `json:"withName,omitempty" firestore:"withName"`
	WithClanTag bool `json:"withClanTag,omitempty" firestore:"withClanTag"`
	WithPins    bool `json:"withPins,omitempty" firestore:"withPins"`
}

type OverviewOptions struct {
	Include          bool     `json:"include,omitempty" firestore:"include"`
	WithTitle        bool     `json:"withTitle,omitempty" firestore:"withTitle"`
	WithLabels       bool     `json:"withLabels,omitempty" firestore:"withLabels"`
	WithAllTimeStats bool     `json:"withAllTimeStats,omitempty" firestore:"withAllTimeStats"`
	Blocks           []string `json:"blocks,omitempty" firestore:"blocks"`
	Type             string   `json:"type,omitempty" firestore:"type"`
}

type VehicleOptions struct {
	Include          bool     `json:"include,omitempty" firestore:"include"`
	Limit            int      `json:"limit,omitempty" firestore:"limit"`
	WithVehicleTier  bool     `json:"withVehicleTier,omitempty" firestore:"withVehicleTier"`
	WithVehicleName  bool     `json:"withVehicleName,omitempty" firestore:"withVehicleName"`
	WithAllTimeStats bool     `json:"withAllTimeStats,omitempty" firestore:"withAllTimeStats"`
	WithLabels       bool     `json:"withLabels,omitempty" firestore:"withLabels"`
	Offset           int      `json:"offset,omitempty" firestore:"offset"`
	Blocks           []string `json:"blocks,omitempty" firestore:"blocks"`
}

const (
	OverviewTypeRegular = "regular"
	OverviewTypeRating  = "rating"
)
