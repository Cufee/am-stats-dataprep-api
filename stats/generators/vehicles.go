package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/logs"
	api "byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	prepTypes "byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateVehiclesCards(stats *api.PlayerRawStats, options types.VehicleOptions, localizer *i18n.Localizer) ([]types.StatsCard, error) {
	var cards []types.StatsCard

	var vehicles []api.VehicleStats
	if len(stats.SessionStats.Vehicles) < options.Offset {
		return cards, fmt.Errorf("generateVehiclesCards: offset %d is greater than session stats vehicles length %d", options.Offset, len(stats.SessionStats.Vehicles))
	} else if len(stats.SessionStats.Vehicles) < (options.Offset + options.Limit) {
		vehicles = stats.SessionStats.Vehicles[options.Offset:]
	} else {
		vehicles = stats.SessionStats.Vehicles[options.Offset : options.Offset+options.Limit]
	}

	for i, vehicle := range vehicles {
		if i >= options.Limit {
			break
		}

		card, err := generateSingleVehicleCard(stats, options, &vehicle, localizer)
		if err != nil {
			logs.Error("Error generating vehicle card for %v: %v", stats.PlayerDetails.ID, err)
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func generateSingleVehicleCard(stats *api.PlayerRawStats, options types.VehicleOptions, vehicle *api.VehicleStats, localizer *i18n.Localizer) (types.StatsCard, error) {
	var row types.StatsCardRow
	for _, block := range options.Blocks {
		if block.GenerationTag == types.BlockWN8Rating.GenerationTag {
			var input prepTypes.DataprepInput
			input.Options.WithAllTime = false // There is no all time WN8 rating for vehicles
			input.Options.WithLabel = options.WithLabels
			input.Options.WithIcons = options.WithIcons
			input.Options.Block = block
			input.Localizer = localizer

			block, err := dataprep.WN8RatingBlock(input, vehicle.TankWN8, 0)
			if err != nil {
				logs.Warning("generateRatingOverviewCard: error generating rating block for %v: %s", stats.PlayerDetails.ID, err)
				continue
			}
			row.Blocks = append(row.Blocks, block)
			continue
		}

		var input prepTypes.DataprepInput
		input.Stats.Session = vehicle.StatsFrame
		allTime, ok := stats.LastSession.Vehicles[fmt.Sprint(vehicle.TankID)]
		if ok {
			input.Stats.AllTime = allTime.StatsFrame
		}
		input.Options.WithAllTime = options.WithAllTimeStats
		input.Options.WithLabel = options.WithLabels
		input.Options.WithIcons = options.WithIcons
		input.Options.Block = block
		input.Localizer = localizer

		block, err := dataprep.BlockFromStats(input)
		if err != nil {
			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", stats.PlayerDetails.ID, err)
			continue
		}
		row.Blocks = append(row.Blocks, block)
	}

	var card types.StatsCard
	if options.WithVehicleName || options.WithVehicleTier {
		var content []types.StatsBlockRowContent
		if options.WithVehicleTier {
			content = append(content, types.StatsBlockRowContent{
				Content: intToRoman(vehicle.TankTier),
				Tags:    []string{utils.TagVehicleTier},
			})
		}
		if options.WithVehicleName {
			content = append(content, types.StatsBlockRowContent{
				Content: vehicle.TankName,
				Tags:    []string{utils.TagVehicleName},
			})
		}

		labelRow := types.StatsCardRow{
			Blocks: []types.StatsBlock{
				{
					Rows: []types.StatsBlockRow{
						{
							Content: content,
						},
					},
					Tags: []string{utils.TagLabel},
				},
			},
		}
		card.Rows = append(card.Rows, labelRow)
	}
	card.Rows = append(card.Rows, row)
	return card, nil
}

func intToRoman(i int) string {
	switch i {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 10:
		return "X"
	default:
		return fmt.Sprint(i)
	}
}
