package presets

import "byvko.dev/repo/am-stats-dataprep-api/stats/types"

var MinimalOptions = types.Options{
	AccountStatus: types.StatusIconsOptions{
		Include: false,
	},
	Notifications: types.NotificationsOptions{
		Include: false,
	},
	Challenges: types.ChallengesOptions{
		Include: false,
	},
	Player: types.PlayerOptions{
		Include: false,
	},
	RatingBattles: types.OverviewOptions{
		Include:    true,
		WithLabels: true,
		Type:       types.OverviewTypeRating,
		Blocks:     []types.BlockOptions{types.BlockBattles, types.BlockWinrate, types.BlockAverageDamage, types.BlockShotAccuracy},
	},
	RegularBattles: types.OverviewOptions{
		Include:    true,
		WithLabels: true,
		Type:       types.OverviewTypeRegular,
		Blocks:     []types.BlockOptions{types.BlockBattles, types.BlockWinrate, types.BlockAverageDamage, types.BlockShotAccuracy},
	},
	VehiclesFull: types.VehicleOptions{
		Include: false,
	},
	VehiclesSlim: types.VehicleOptions{
		Include: false,
	},
}
