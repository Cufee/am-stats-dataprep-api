package presets

import "byvko.dev/repo/am-stats-dataprep-api/stats/types"

const (
	DetailedVehiclesLimit = 3
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
		Blocks:  []types.BlockOptions{},
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
		WithIcons:        true,
		WithAllTimeStats: true,
		Type:             types.OverviewTypeRating,
		Blocks:           []types.BlockOptions{types.BlockBattles, types.BlockAverageDamage.WithIcon(), types.BlockWinrate.WithIcon(), types.BlockShotAccuracy.WithIcon()},
	},
	RegularBattles: types.OverviewOptions{
		Include:          true,
		WithTitle:        true,
		WithLabels:       true,
		WithIcons:        true,
		WithAllTimeStats: true,
		Type:             types.OverviewTypeRegular,
		Blocks:           []types.BlockOptions{types.BlockBattles, types.BlockAverageDamage.WithIcon(), types.BlockWinrate.WithIcon(), types.BlockShotAccuracy.WithIcon(), types.BlockWN8Rating.WithIcon()},
	},
	VehiclesFull: types.VehicleOptions{
		Include:          true,
		Limit:            3,
		WithVehicleTier:  true,
		WithVehicleName:  true,
		WithIcons:        true,
		WithAllTimeStats: true,
		WithLabels:       true,
		Blocks:           []types.BlockOptions{types.BlockBattles, types.BlockAverageDamage.WithIcon(), types.BlockWinrate.WithIcon(), types.BlockShotAccuracy.WithIcon(), types.BlockWN8Rating.WithIcon()},
	},
	VehiclesSlim: types.VehicleOptions{
		Include:          true,
		Limit:            3,
		WithVehicleTier:  true,
		WithIcons:        true,
		WithVehicleName:  true,
		WithAllTimeStats: false,
		WithLabels:       true,
		Blocks:           []types.BlockOptions{types.BlockAverageDamage, types.BlockWinrateWithBattles.WithIcon(), types.BlockShotAccuracy.WithIcon(), types.BlockWN8Rating.WithIcon()},
	},
}
