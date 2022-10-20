package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-core/stats/ratings/wn8/v1"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/stats/v3"
)

func GenerateVehiclesCards(layout *logic.CardLayout, layoutName string, session []stats.VehicleStats, allTime map[int]stats.VehicleStats, printer func(string) string) []block.Block {
	cards := make([]block.Block, 0, len(session))
	for _, vehicle := range session {
		allTimeVehicle := allTime[vehicle.TankID]
		card := generateSingleVehicleCard(layout, layoutName, vehicle, allTimeVehicle, printer)
		if card != nil {
			cards = append(cards, *card)
		}
	}
	return cards
}

func generateSingleVehicleCard(layout *logic.CardLayout, layoutName string, session, allTime stats.VehicleStats, printer func(string) string) *block.Block {
	var card block.Block
	card.Style = layout.CardStyle
	card.ContentType = block.ContentTypeBlocks
	layout.Title.String = fmt.Sprintf("%s %s", intToRoman(session.TankTier), session.TankName)
	var cardRows []block.Block
	cardRows = append(cardRows, layout.Title.ToBlock(nil))

	var contentBlock block.Block
	contentBlock.Style = layout.ContentStyle
	contentBlock.ContentType = block.ContentTypeBlocks
	content := make([]block.Block, 0, len(layout.Blocks))

	for _, definition := range layout.Blocks {
		switch definition.ValueKind {
		case logic.WN8OverOne:
			b := WN8BlockFromStats(layoutName, definition, session.Ratings[wn8.WN8], -1, printer)
			if b != nil {
				content = append(content, *b)
			}
		default:
			b := BlockFromStats(layoutName, definition, session.Stats, allTime.Stats, printer)
			if b != nil {
				content = append(content, *b)
			}
		}
	}
	if len(content) == 0 {
		return nil
	}

	contentBlock.Content = content
	cardRows = append(cardRows, contentBlock)
	card.Content = cardRows
	return &card
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
		return ""
	}
}
