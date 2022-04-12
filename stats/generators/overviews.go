package generators

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/stats/v1"
)

func GenerateOverviewCard(layout *logic.CardLayout, layoutName string, data *stats.PlayerRawStats, printer func(string) string) *block.Block {
	var card block.Block
	card.Style = layout.CardStyle
	card.ContentType = block.ContentTypeBlocks
	layout.Title.Printer = printer
	var cardRows []block.Block
	cardRows = append(cardRows, layout.Title.ToBlock(nil))

	var contentBlock block.Block
	contentBlock.Style = layout.ContentStyle
	contentBlock.ContentType = block.ContentTypeBlocks
	content := make([]block.Block, 0, len(layout.Blocks))
	for _, b := range layout.Blocks {
		switch b.ValueKind {
		case logic.WN8OverOne:
			block := WN8BlockFromStats(layoutName, b, data.SessionStats.SessionRating, data.PlayerDetails.CareerWN8, printer)
			if block != nil {
				content = append(content, *block)
			}

		default:
			block := BlockFromStats(layoutName, b, data.SessionStats.StatsAll, data.PlayerDetails.Stats.All, printer)
			if block != nil {
				content = append(content, *block)
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
