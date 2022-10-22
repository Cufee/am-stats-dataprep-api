package stats

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/localization"
	"byvko.dev/repo/am-stats-dataprep-api/stats/generators"
	"byvko.dev/repo/am-stats-dataprep-api/stats/helpers"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"

	api "github.com/byvko-dev/am-types/api/stats/v1"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	types "github.com/byvko-dev/am-types/stats/v3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type StatsResponse struct {
	// StatusIcons block.Block `json:"statusIcons" bson:"statusIcons"`
	Cards       block.Block `json:"cards" bson:"cards"`
	FailedCards []string    `json:"failedCards" bson:"failedCards"`
	LastBattle  int         `json:"lastBattle" bson:"lastBattle"`
	Style       string      `json:"style" bson:"style"`
}

func CompilePlayerStatsCards(stats *api.ResponsePayload, options *logic.LayoutOptions, locale, styleName string) (StatsResponse, error) {
	if stats == nil {
		return StatsResponse{}, fmt.Errorf("stats is nil")
	}

	// Localization
	localizer := localization.InitLocalizer(localization.LocaleStringFromLanguage(locale))
	printer := func(s string) string {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: s,
		})
		return label
	}

	var response StatsResponse
	var cards []block.Block

	if options.PlayerInfo != nil {
		card := generators.GeneratePlayerCard(options.PlayerInfo, options.LayoutName, stats, printer)
		if card != nil {
			cards = append(cards, *card)
		}
	}

	if options.RatingOverview != nil && stats.Session.Rating.Total.Battles > 0 {
		overview := generators.GenerateRatingOverviewCard(options.RatingOverview, options.LayoutName, stats, printer)
		if overview != nil {
			cards = append(cards, *overview)
		}
	}

	if options.RandomOverview != nil && stats.Session.Regular.Total.Battles > 0 {
		overview := generators.GenerateRandomOverviewCard(options.RandomOverview, options.LayoutName, stats, printer)
		if overview != nil {
			cards = append(cards, *overview)
		}
	}

	var allVehicles []types.VehicleStats
	for _, v := range stats.Session.Regular.Vehicles {
		if v.Stats.Battles > 0 {
			allVehicles = append(allVehicles, v)
		}
	}
	allVehicles = helpers.SortTanks(allVehicles, options.VehiclesSort)

	var slimVehiclesOffset int
	if options.VehiclesFullOverview != nil && len(stats.Session.Regular.Vehicles) > 0 {
		vehicles := allVehicles
		if len(vehicles) > options.VehiclesFullOverview.Limit {
			vehicles = vehicles[:options.VehiclesFullOverview.Limit]
		}

		vehiclesFull := generators.GenerateVehiclesCards(options.VehiclesFullOverview, options.LayoutName, vehicles, stats.Snapshot.Regular.Vehicles, locale, printer)
		cards = append(cards, vehiclesFull...)
		slimVehiclesOffset = len(vehiclesFull)
	}

	if options.VehiclesSlimOverview != nil && len(stats.Session.Regular.Vehicles) > slimVehiclesOffset {
		vehicles := allVehicles[slimVehiclesOffset:]
		if len(vehicles) > options.VehiclesSlimOverview.Limit {
			vehicles = vehicles[:options.VehiclesSlimOverview.Limit]
		}

		vehiclesCards := generators.GenerateVehiclesCards(options.VehiclesSlimOverview, options.LayoutName, vehicles, stats.Snapshot.Regular.Vehicles, locale, printer)
		cards = append(cards, vehiclesCards...)
	}

	cardBlock := block.Block{
		Content:     cards,
		ContentType: block.ContentTypeBlocks,
		Style:       shared.DefaultFont.Merge(options.WrapperStyle),
	}

	response.Cards = cardBlock
	response.LastBattle = int(stats.Timestamp.Unix())
	return response, nil
}
