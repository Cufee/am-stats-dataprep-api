package stats

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/localization"
	"byvko.dev/repo/am-stats-dataprep-api/stats/generators"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"
	types "byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
	api "github.com/byvko-dev/am-types/stats/v1"

	"github.com/byvko-dev/am-core/logs"
)

func CompilePlayerStatsCards(stats *api.PlayerRawStats, options settings.Options, styleName string) (types.StatsResponse, error) {
	if stats == nil {
		return types.StatsResponse{}, fmt.Errorf("stats is nil")
	}

	// Localization
	localizer := localization.InitLocalizer(localization.LocaleStringFromLanguage(options.Locale))

	var response types.StatsResponse
	var cards []block.Block

	if options.AccountStatus.Include {
		statusIcons, err := generators.GenerateStatusIcons(stats, options.AccountStatus)
		if err != nil {
			logs.Error("Failed to generate status icons for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.AccountStatus")
		} else {
			statusBlock := block.Block{
				ContentType: block.ContentTypeBlocks,
				Content:     statusIcons,
			}
			response.StatusIcons = statusBlock
		}
	}

	if options.Notifications.Include {
		notifications, err := generators.GenerateNotificationsCards(stats, options.Notifications)
		if err != nil {
			logs.Error("Failed to generate notifications for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.Notifications")
		} else {
			cards = append(cards, notifications...)
		}
	}

	if options.Challenges.Include {
		challenges, err := generators.GenerateChallengesCards(stats, options.Challenges)
		if err != nil {
			logs.Error("Failed to generate challenges for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.Challenges")
		} else {
			cards = append(cards, challenges...)
		}
	}

	if options.Player.Include {
		playerCard, err := generators.GeneratePlayerCard(stats, options.Player, styleName)
		if err != nil {
			logs.Error("Failed to generate player card for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.Player")
		} else {
			cards = append(cards, playerCard)
		}
	}

	if options.RatingBattles.Include && stats.SessionStats.BattlesRating > 0 {
		ratingBattles, err := generators.GenerateOverviewCard(stats, options.RatingBattles, localizer, styleName)
		if err != nil {
			logs.Error("Failed to generate rating battles for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.RatingBattles")
		} else {
			cards = append(cards, ratingBattles)
		}
	}

	if options.RegularBattles.Include && stats.SessionStats.BattlesAll > 0 {
		regularBattles, err := generators.GenerateOverviewCard(stats, options.RegularBattles, localizer, styleName)
		if err != nil {
			logs.Error("Failed to generate regular battles for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.RatingBattles")
		} else {
			cards = append(cards, regularBattles)
		}
	}

	var slimVehiclesOffset int = 0
	if options.VehiclesFull.Include && len(stats.SessionStats.Vehicles) > 0 {
		vehiclesFull, err := generators.GenerateVehiclesCards(stats, options.VehiclesFull, localizer, styleName)
		if err != nil {
			logs.Error("Failed to generate vehicles full for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.VehiclesFull")
		} else {
			cards = append(cards, vehiclesFull...)
			slimVehiclesOffset = len(vehiclesFull)
		}
	}

	if options.VehiclesSlim.Include && len(stats.SessionStats.Vehicles) >= slimVehiclesOffset {
		options.VehiclesSlim.Offset = slimVehiclesOffset
		vehiclesSlim, err := generators.GenerateVehiclesCards(stats, options.VehiclesSlim, localizer, styleName)
		if err != nil {
			logs.Error("Failed to generate vehicles slim for %v: %v", stats.PlayerDetails.ID, err)
			response.FailedCards = append(response.FailedCards, "options.VehiclesSlim")
		} else {
			cards = append(cards, vehiclesSlim...)
		}
	}

	bgStyle := styles.LoadBackground(styleName, "wrapper")
	cardBlock := block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     cards,
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "wrapper")).Merge(bgStyle),
		Tags:        []string{"wrapper"},
	}

	response.Cards = cardBlock
	response.LastBattle = stats.PlayerDetails.LastBattle
	return response, nil
}
