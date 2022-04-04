package generators

import (
	"fmt"
	"sync"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	prepTypes "byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/dataprep/v1/block"
	"github.com/byvko-dev/am-types/dataprep/v1/settings"
	api "github.com/byvko-dev/am-types/stats/v1"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type cardWithPosition struct {
	card  block.Block
	index int
}

func GenerateVehiclesCards(stats *api.PlayerRawStats, options settings.VehicleOptions, localizer *i18n.Localizer) ([]block.Block, error) {
	var vehicles []api.VehicleStats
	if len(stats.SessionStats.Vehicles) < options.Offset {
		return nil, fmt.Errorf("generateVehiclesCards: offset %d is greater than session stats vehicles length %d", options.Offset, len(stats.SessionStats.Vehicles))
	} else if len(stats.SessionStats.Vehicles) < (options.Offset + options.Limit) {
		vehicles = stats.SessionStats.Vehicles[options.Offset:]
	} else {
		vehicles = stats.SessionStats.Vehicles[options.Offset : options.Offset+options.Limit]
	}

	wg := new(sync.WaitGroup)
	cardsChan := make(chan cardWithPosition, len(vehicles))

	for i, vehicle := range vehicles {
		wg.Add(1)
		go func(i int, vehicle api.VehicleStats) {
			defer wg.Done()
			card, err := generateSingleVehicleCard(stats, options, &vehicle, localizer)
			if err != nil {
				logs.Error("Error generating vehicle card for %v: %v", stats.PlayerDetails.ID, err)
			}
			cardsChan <- cardWithPosition{
				card:  card,
				index: i,
			}
		}(i, vehicle)
	}
	wg.Wait()
	close(cardsChan)

	cards := make([]block.Block, len(vehicles))
	for card := range cardsChan {
		cards[card.index] = card.card
	}

	return cards, nil
}

func generateSingleVehicleCard(stats *api.PlayerRawStats, options settings.VehicleOptions, vehicle *api.VehicleStats, localizer *i18n.Localizer) (block.Block, error) {
	var rowContent []block.Block
	for _, b := range options.Blocks {
		if b.GenerationTag == types.BlockWN8Rating.GenerationTag {
			var input prepTypes.DataprepInput
			input.Options.WithAllTime = false // There is no all time WN8 rating for vehicles
			input.Options.WithLabel = options.WithLabels
			input.Options.WithIcons = options.WithIcons
			input.Options.Block = b
			input.Localizer = localizer

			block, err := dataprep.WN8RatingBlock(input, vehicle.TankWN8, 0)
			if err != nil {
				logs.Warning("generateRatingOverviewCard: error generating rating block for %v: %s", stats.PlayerDetails.ID, err)
				continue
			}
			rowContent = append(rowContent, block)
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
		input.Options.Block = b
		input.Localizer = localizer

		block, err := dataprep.BlockFromStats(input)
		if err != nil {
			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", stats.PlayerDetails.ID, err)
			continue
		}
		rowContent = append(rowContent, block)
	}

	var cardRows []block.Block
	if options.WithVehicleName || options.WithVehicleTier {
		var content []block.Block
		if options.WithVehicleTier {
			content = append(content, block.Block{
				ContentType: block.ContentTypeText,
				Content:     intToRoman(vehicle.TankTier),
				Tags:        []string{utils.TagVehicleTier},
			})
		}
		if options.WithVehicleName {
			content = append(content, block.Block{
				ContentType: block.ContentTypeText,
				Content:     vehicle.TankName,
				Tags:        []string{utils.TagVehicleName},
			})
		}

		labelRow := block.Block{
			ContentType: block.ContentTypeBlocks,
			Content: []block.Block{
				{
					ContentType: block.ContentTypeBlocks,
					Content: []block.Block{
						{
							ContentType: block.ContentTypeBlocks,
							Content:     content,
						},
					},
					Tags: []string{utils.TagLabel},
				},
			},
			Tags: []string{"title_row"},
		}
		cardRows = append(cardRows, labelRow)
	}
	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
	})
	return block.Block{
		ContentType: block.ContentTypeBlocks,
		Style: block.Style{
			AlignItems: block.AlignItemsVertical,
		},
		Content: cardRows,
	}, nil
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
