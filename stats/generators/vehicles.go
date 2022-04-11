package generators

import (
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

type cardWithPosition struct {
	card  block.Block
	index int
}

// func GenerateVehiclesCards(stats *api.PlayerRawStats, options settings.VehicleOptions, localizer *i18n.Localizer, styleName string) ([]block.Block, error) {
// 	var vehicles []api.VehicleStats
// 	if len(stats.SessionStats.Vehicles) < options.Offset {
// 		return nil, fmt.Errorf("generateVehiclesCards: offset %d is greater than session stats vehicles length %d", options.Offset, len(stats.SessionStats.Vehicles))
// 	} else if len(stats.SessionStats.Vehicles) < (options.Offset + options.Limit) {
// 		vehicles = stats.SessionStats.Vehicles[options.Offset:]
// 	} else {
// 		vehicles = stats.SessionStats.Vehicles[options.Offset : options.Offset+options.Limit]
// 	}

// 	wg := new(sync.WaitGroup)
// 	cardsChan := make(chan cardWithPosition, len(vehicles))

// 	for i, vehicle := range vehicles {
// 		wg.Add(1)
// 		go func(i int, vehicle api.VehicleStats) {
// 			defer wg.Done()
// 			card, err := generateSingleVehicleCard(stats, options, &vehicle, localizer, styleName)
// 			if err != nil {
// 				logs.Error("Error generating vehicle card for %v: %v", stats.PlayerDetails.ID, err)
// 			}
// 			cardsChan <- cardWithPosition{
// 				card:  card,
// 				index: i,
// 			}
// 		}(i, vehicle)
// 	}
// 	wg.Wait()
// 	close(cardsChan)

// 	cards := make([]block.Block, len(vehicles))
// 	for card := range cardsChan {
// 		cards[card.index] = block.Block{
// 			ContentType: block.ContentTypeBlocks,
// 			Content:     []block.Block{card.card},
// 			Style:       styles.LoadWithTags(styleName, "card"),
// 		}
// 	}

// 	return cards, nil
// }

// func generateSingleVehicleCard(stats *api.PlayerRawStats, options settings.VehicleOptions, vehicle *api.VehicleStats, localizer *i18n.Localizer, styleName string) (block.Block, error) {
// 	var rowContent []block.Block
// 	for _, b := range options.Blocks {
// 		if b.GenerationTag == tags.GenerationTagWN8Rating {
// 			var input prepTypes.DataprepInput
// 			input.Options.WithAllTime = false // There is no all time WN8 rating for vehicles
// 			input.Options.WithLabel = options.WithLabels
// 			input.Options.WithIcons = options.WithIcons
// 			input.Options.Style = styleName
// 			input.Options.Block = b
// 			input.Localizer = localizer
// 			if vehicle.TankWN8 < 0 {
// 				continue
// 			}

// 			block, err := dataprep.WN8RatingBlock(input, vehicle.TankWN8, 0)
// 			if err != nil {
// 				logs.Warning("generateRatingOverviewCard: error generating rating block for %v: %s", stats.PlayerDetails.ID, err)
// 				continue
// 			}
// 			rowContent = append(rowContent, block)
// 			continue
// 		}

// 		var input prepTypes.DataprepInput
// 		input.Stats.Session = vehicle.StatsFrame
// 		allTime, ok := stats.LastSession.Vehicles[fmt.Sprint(vehicle.TankID)]
// 		if ok {
// 			input.Stats.AllTime = allTime.StatsFrame
// 		}
// 		input.Options.WithAllTime = options.WithAllTimeStats
// 		input.Options.WithLabel = options.WithLabels
// 		input.Options.WithIcons = options.WithIcons
// 		input.Options.Style = styleName
// 		input.Options.Block = b
// 		input.Localizer = localizer

// 		block, err := dataprep.BlockFromStats(input)
// 		if err != nil {
// 			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", stats.PlayerDetails.ID, err)
// 			continue
// 		}
// 		rowContent = append(rowContent, block)
// 	}

// 	var cardRows []block.Block
// 	if options.WithVehicleName || options.WithVehicleTier {
// 		var content []block.Block
// 		if options.WithVehicleTier {
// 			content = append(content, block.Block{
// 				ContentType: block.ContentTypeText,
// 				Content:     intToRoman(vehicle.TankTier),
// 				Style:       styles.LoadWithTags(styleName, utils.TagVehicleTier),
// 			})
// 		}
// 		if options.WithVehicleName {
// 			content = append(content, block.Block{
// 				ContentType: block.ContentTypeText,
// 				Content:     vehicle.TankName,
// 				Style:       styles.LoadWithTags(styleName, utils.TagVehicleName),
// 			})
// 		}

// 		labelRow := block.Block{
// 			ContentType: block.ContentTypeBlocks,
// 			Content: []block.Block{
// 				{
// 					ContentType: block.ContentTypeBlocks,
// 					Content:     content,
// 					Style:       styles.LoadWithTags(styleName, utils.TagLabel),
// 				},
// 			},
// 			Style: styles.LoadWithTags(styleName, "titleRow"),
// 		}
// 		cardRows = append(cardRows, labelRow)
// 	}

// 	// Find fixtag
// 	fixTag := "fixIcon-false"
// 	for _, b := range rowContent {
// 		if slices.Contains(b.Tags, "fixIcon-true") > -1 {
// 			fixTag = "fixIcon-true"
// 			break
// 		}
// 	}

// 	cardRows = append(cardRows, block.Block{
// 		ContentType: block.ContentTypeBlocks,
// 		Content:     rowContent,
// 		Style:       styles.LoadWithTags(styleName, "vehicleRowContent", "statsContent", fixTag),
// 	})
// 	return block.Block{
// 		ContentType: block.ContentTypeBlocks,
// 		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "vehicleOverview")),
// 		Content:     cardRows,
// 	}, nil
// }

// func intToRoman(i int) string {
// 	switch i {
// 	case 1:
// 		return "I"
// 	case 2:
// 		return "II"
// 	case 3:
// 		return "III"
// 	case 4:
// 		return "IV"
// 	case 5:
// 		return "V"
// 	case 6:
// 		return "VI"
// 	case 7:
// 		return "VII"
// 	case 8:
// 		return "VIII"
// 	case 9:
// 		return "IX"
// 	case 10:
// 		return "X"
// 	default:
// 		return ""
// 	}
// }
