package stats

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/localization"
	"byvko.dev/repo/am-stats-dataprep-api/logs"
	api "byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/generators"
	types "byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func CompilePlayerStatsCards(stats *api.PlayerRawStats, options types.Options) (types.StatsResponse, error) {
	if stats == nil {
		return types.StatsResponse{}, fmt.Errorf("stats is nil")
	}

	// Localization
	localizer := localization.InitLocalizer(localization.LocaleStringFromLanguage(options.Locale))

	var response types.StatsResponse
	var cards []types.StatsCard

	if options.AccountStatus.Include {
		statusIcons, err := generators.GenerateStatusIcons(stats, options.AccountStatus)
		if err != nil {
			logs.Error("Failed to generate status icons for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.AccountStatus")
		}
		response.StatusIcons = statusIcons
	}

	if options.Notifications.Include {
		notifications, err := generators.GenerateNotificationsCards(stats, options.Notifications)
		if err != nil {
			logs.Error("Failed to generate notifications for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.Notifications")
		}
		cards = append(cards, notifications...)
	}

	if options.Challenges.Include {
		challenges, err := generators.GenerateChallengesCards(stats, options.Challenges)
		if err != nil {
			logs.Error("Failed to generate challenges for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.Challenges")
		}
		cards = append(cards, challenges...)
	}

	if options.Player.Include {
		playerCard, err := generators.GeneratePlayerCard(stats, options.Player)
		if err != nil {
			logs.Error("Failed to generate player card for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.Player")
		}
		cards = append(cards, playerCard)
	}

	if options.RatingBattles.Include && stats.SessionStats.BattlesRating > 0 {
		ratingBattles, err := generators.GenerateOverviewCard(stats, options.RatingBattles, localizer)
		if err != nil {
			logs.Error("Failed to generate rating battles for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.RatingBattles")
		}
		cards = append(cards, ratingBattles)
	}

	if options.RegularBattles.Include && stats.SessionStats.BattlesAll > 0 {
		regularBattles, err := generators.GenerateOverviewCard(stats, options.RegularBattles, localizer)
		if err != nil {
			logs.Error("Failed to generate regular battles for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.RatingBattles")
		}
		cards = append(cards, regularBattles)
	}

	var slimVehiclesOffset int = 0
	if options.VehiclesFull.Include && len(stats.SessionStats.Vehicles) > 0 {
		vehiclesFull, err := generators.GenerateVehiclesCards(stats, options.VehiclesFull, localizer)
		if err != nil {
			logs.Error("Failed to generate vehicles full for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.VehiclesFull")
		}
		cards = append(cards, vehiclesFull...)
		slimVehiclesOffset = len(vehiclesFull)
	}

	if options.VehiclesSlim.Include && len(stats.SessionStats.Vehicles) >= slimVehiclesOffset {
		options.VehiclesSlim.Offset = slimVehiclesOffset
		vehiclesSlim, err := generators.GenerateVehiclesCards(stats, options.VehiclesSlim, localizer)
		if err != nil {
			logs.Error("Failed to generate vehicles slim for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.VehiclesSlim")
		}
		cards = append(cards, vehiclesSlim...)
	}

	response.Cards = cards
	response.LastBattle = stats.PlayerDetails.LastBattle
	if len(response.Cards) == 0 {
		logs.Error("Failed to generate any cards for %v", stats.PlayerDetails.ID)
		return response, fmt.Errorf("failed to generate any cards for %v", stats.PlayerDetails.ID)
	}
	return response, nil
}
