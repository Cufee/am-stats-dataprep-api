package types

type Options struct {
	// Status icons
	AccountStatus StatusIconsOptions

	// Notifications
	Notifications NotificationsOptions

	// Challenges
	Challenges ChallengesOptions

	// Player info and general stats
	Player         PlayerOptions
	RatingBattles  OverviewOptions
	RegularBattles OverviewOptions

	// Per vehicle stats
	VehiclesFull VehicleOptions
	VehiclesSlim VehicleOptions
}

type StatusIconsOptions struct {
	Include bool
	Limit   int
}

type NotificationsOptions struct {
	Include bool
	Blocks  []string
}

type ChallengesOptions struct {
	Include bool
	Limit   int
	Blocks  []string
	Types   []string
}

type PlayerOptions struct {
	Include     bool
	WithName    bool
	WithClanTag bool
}

type OverviewOptions struct {
	Include          bool
	WithLabels       bool
	WithAllTimeStats bool
	Blocks           []string
}

type VehicleOptions struct {
	Include          bool
	Limit            int
	WithVehicleTier  bool
	WithVehicleName  bool
	WithAllTimeStats bool
	WithLabels       bool
	Blocks           []string
}
