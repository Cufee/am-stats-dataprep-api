package presets

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-types/dataprep/v1/settings"
)

var MinimalOptions = settings.Options{
	AccountStatus: settings.StatusIconsOptions{
		Include: false,
	},
	Notifications: settings.NotificationsOptions{
		Include: false,
	},
	Challenges: settings.ChallengesOptions{
		Include: false,
	},
	Player: settings.PlayerOptions{
		Include: false,
	},
	RatingBattles: settings.OverviewOptions{
		Include:    true,
		WithLabels: true,
		Type:       settings.OverviewTypeRating,
		Blocks:     []settings.BlockOptions{types.BlockBattles, types.BlockWinrate, types.BlockAverageDamage, types.BlockShotAccuracy},
	},
	RegularBattles: settings.OverviewOptions{
		Include:    true,
		WithLabels: true,
		Type:       settings.OverviewTypeRegular,
		Blocks:     []settings.BlockOptions{types.BlockBattles, types.BlockWinrate, types.BlockAverageDamage, types.BlockShotAccuracy},
	},
	VehiclesFull: settings.VehicleOptions{
		Include: false,
	},
	VehiclesSlim: settings.VehicleOptions{
		Include: false,
	},
}
