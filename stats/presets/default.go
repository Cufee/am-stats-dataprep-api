package presets

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
)

const (
	defaultDetailedVehiclesLimit = 3
	defaultSlimVehiclesLimit     = 4
)

var DefaultOptions = settings.Options{
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
		Blocks:           []settings.BlockOptions{types.BlockBattles, types.BlockAverageDamage.WithIcon().WithInvisibleIcon(), types.BlockWinrate.WithIcon().WithInvisibleIcon(), types.BlockShotAccuracy.WithIcon().WithInvisibleIcon()},
	},
	RegularBattles: settings.OverviewOptions{
		Include:          true,
		WithTitle:        true,
		WithLabels:       true,
		WithIcons:        true,
		WithAllTimeStats: true,
		Type:             settings.OverviewTypeRegular,
		Blocks:           []settings.BlockOptions{types.BlockBattles, types.BlockAverageDamage.WithIcon().WithInvisibleIcon(), types.BlockWinrate.WithIcon().WithInvisibleIcon(), types.BlockShotAccuracy.WithIcon().WithInvisibleIcon(), types.BlockWN8Rating.WithIcon().WithInvisibleIcon()},
	},
	VehiclesFull: settings.VehicleOptions{
		Include:          true,
		Limit:            defaultDetailedVehiclesLimit,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithIcons:        true,
		WithAllTimeStats: true,
		WithLabels:       true,
		Blocks:           []settings.BlockOptions{types.BlockBattles, types.BlockAverageDamage.WithIcon().WithInvisibleIcon(), types.BlockWinrate.WithIcon().WithInvisibleIcon(), types.BlockShotAccuracy.WithIcon().WithInvisibleIcon(), types.BlockWN8Rating.WithIcon().WithInvisibleIcon()},
	},
	VehiclesSlim: settings.VehicleOptions{
		Limit:            defaultSlimVehiclesLimit,
		Include:          true,
		WithVehicleTier:  true,
		WithIcons:        true,
		WithVehicleName:  true,
		WithAllTimeStats: false,
		WithLabels:       true,
		Blocks:           []settings.BlockOptions{types.BlockAverageDamage, types.BlockWinrateWithBattles.WithIcon().WithInvisibleIcon(), types.BlockShotAccuracy.WithIcon().WithInvisibleIcon(), types.BlockWN8Rating.WithIcon().WithInvisibleIcon()},
	},
}
