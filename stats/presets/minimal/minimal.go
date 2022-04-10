package minimal

import (
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
)

var Options = settings.Options{
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
		Blocks:     []settings.BlockOptions{BlockBattles, BlockWinrate, BlockAverageDamage, BlockShotAccuracy},
	},
	RegularBattles: settings.OverviewOptions{
		Include:    true,
		WithLabels: true,
		Type:       settings.OverviewTypeRegular,
		Blocks:     []settings.BlockOptions{BlockBattles, BlockWinrate, BlockAverageDamage, BlockShotAccuracy},
	},
	VehiclesFull: settings.VehicleOptions{
		Include: false,
	},
	VehiclesSlim: settings.VehicleOptions{
		Include: false,
	},
}
