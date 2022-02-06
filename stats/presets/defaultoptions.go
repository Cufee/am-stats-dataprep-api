package presets

import "byvko.dev/repo/am-stats-dataprep-api/stats/types"

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
		Blocks:  []string{types.BlockChallengeSource, types.BlockChallengeName, types.BlockChallengeProgress, types.BlockChallengeTimeLeft},
		Types:   []string{types.ChallengeTypeAll},
	},
	Player: types.PlayerOptions{
		Include:     true,
		WithName:    true,
		WithClanTag: true,
	},
	RatingBattles: types.OverviewOptions{
		Include:          true,
		WithLabels:       true,
		WithAllTimeStats: true,
		Blocks:           []string{types.BlockBattles, types.BlockAverageDamage, types.BlockWinrate, types.BlockWN8Rating},
	},
	RegularBattles: types.OverviewOptions{
		Include:          true,
		WithLabels:       true,
		WithAllTimeStats: true,
		Blocks:           []string{types.BlockBattles, types.BlockAverageDamage, types.BlockWinrate, types.BlockWN8Rating},
	},
	VehiclesFull: types.VehicleOptions{
		Include:          true,
		Limit:            3,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithAllTimeStats: true,
		WithLabels:       true,
		Blocks:           []string{types.BlockBattles, types.BlockAverageDamage, types.BlockWinrate, types.BlockWN8Rating},
	},
	VehiclesSlim: types.VehicleOptions{
		Include:          true,
		Limit:            3,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithAllTimeStats: false,
		WithLabels:       true,
		Blocks:           []string{types.BlockAverageDamage, types.BlockWinrateWithBattles, types.BlockWN8Rating},
	},
}
