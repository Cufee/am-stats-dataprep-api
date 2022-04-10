package legacy

import (
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
)

const (
	legacyDetailedVehiclesLimit = 3
	legacySlimVehiclesLimit     = 4
)

var Options = settings.Options{
	AccountStatus: settings.StatusIconsOptions{
		Include: false, // Not complete
		Limit:   3,
	},
	Notifications: settings.NotificationsOptions{
		Include: false, // Not complete
	},
	Challenges: settings.ChallengesOptions{
		Include: false, // Not complete
		Limit:   1,
		Blocks:  []settings.BlockOptions{},
		// Blocks:  []string{types.BlockChallengeSource, types.BlockChallengeName, types.BlockChallengeProgress, types.BlockChallengeTimeLeft},
		Types: []string{},
		// Types:   []string{types.ChallengeTypeAll},
	},
	Player: settings.PlayerOptions{
		Include:     true,
		WithName:    true,
		WithClanTag: true,
		WithPins:    false, // Not complete
	},
	RatingBattles: settings.OverviewOptions{
		Include:          true,
		WithTitle:        true,
		WithLabels:       true,
		WithIcons:        true,
		WithAllTimeStats: true,
		Type:             settings.OverviewTypeRating,
		Blocks:           []settings.BlockOptions{BlockBattles, BlockAverageDamage.WithIcon().WithInvisibleIcon(), BlockWinrate.WithIcon().WithInvisibleIcon(), BlockShotAccuracy.WithIcon().WithInvisibleIcon()},
	},
	RegularBattles: settings.OverviewOptions{
		Include:          true,
		WithTitle:        true,
		WithLabels:       true,
		WithIcons:        true,
		WithAllTimeStats: true,
		Type:             settings.OverviewTypeRegular,
		Blocks:           []settings.BlockOptions{BlockBattles, BlockAverageDamage.WithIcon().WithInvisibleIcon(), BlockWinrate.WithIcon().WithInvisibleIcon(), BlockWN8Rating.WithIcon().WithInvisibleIcon()},
	},
	VehiclesFull: settings.VehicleOptions{
		Include:          true,
		Limit:            legacyDetailedVehiclesLimit,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithIcons:        true,
		WithAllTimeStats: true,
		WithLabels:       true,
		Blocks:           []settings.BlockOptions{BlockBattles, BlockAverageDamage.WithIcon().WithInvisibleIcon(), BlockWinrate.WithIcon().WithInvisibleIcon(), BlockWN8Rating.WithIcon().WithInvisibleIcon()},
	},
	VehiclesSlim: settings.VehicleOptions{
		Limit:            legacySlimVehiclesLimit,
		Include:          true,
		WithVehicleTier:  true,
		WithIcons:        true,
		WithVehicleName:  true,
		WithAllTimeStats: false,
		WithLabels:       true,
		Blocks:           []settings.BlockOptions{BlockAverageDamage, BlockWinrateWithBattles.WithIcon().WithInvisibleIcon(), BlockWN8Rating.WithIcon().WithInvisibleIcon()},
	},
}
