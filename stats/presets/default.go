package presets

import "byvko.dev/repo/am-stats-dataprep-api/stats/types"

const (
	DetailsVehiclesLimit = 3
)

var DefaultOptions = types.Options{
	AccountStatus: types.StatusIconsOptions{
		Include: false, // Not complete
		Limit:   3,
	},
	Notifications: types.NotificationsOptions{
		Include: false, // Not complete
	},
	Challenges: types.ChallengesOptions{
		Include: false, // Not complete
		Limit:   1,
		Blocks:  []string{},
		// Blocks:  []string{types.BlockChallengeSource, types.BlockChallengeName, types.BlockChallengeProgress, types.BlockChallengeTimeLeft},
		Types: []string{},
		// Types:   []string{types.ChallengeTypeAll},
	},
	Player: types.PlayerOptions{
		Include:     true,
		WithName:    true,
		WithClanTag: true,
		WithPins:    false, // Not complete
	},
	RatingBattles: types.OverviewOptions{
		Include:          true,
		WithTitle:        true,
		WithLabels:       true,
		WithAllTimeStats: true,
		Type:             types.OverviewTypeRating,
		Blocks:           []string{types.BlockBattles, types.BlockAverageDamage, types.BlockWinrate, types.BlockWN8Rating},
	},
	RegularBattles: types.OverviewOptions{
		Include:          true,
		WithTitle:        true,
		WithLabels:       true,
		WithAllTimeStats: true,
		Type:             types.OverviewTypeRegular,
		Blocks:           []string{types.BlockBattles, types.BlockAverageDamage, types.BlockWinrate, types.BlockWN8Rating},
	},
	VehiclesFull: types.VehicleOptions{
		Include:          true,
		Limit:            DetailsVehiclesLimit,
		Offset:           0,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithAllTimeStats: true,
		WithLabels:       true,
		Blocks:           []string{types.BlockBattles, types.BlockAverageDamage, types.BlockWinrate, types.BlockWN8Rating},
	},
	VehiclesSlim: types.VehicleOptions{
		Include:          true,
		Limit:            3,
		Offset:           DetailsVehiclesLimit,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithAllTimeStats: false,
		WithLabels:       true,
		Blocks:           []string{types.BlockAverageDamage, types.BlockWinrateWithBattles, types.BlockWN8Rating},
	},
}
