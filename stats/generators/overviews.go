package generators

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"github.com/byvko-dev/am-core/stats/ratings/wn8/v1"
	"github.com/byvko-dev/am-types/api/stats/v1"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func GenerateRatingOverviewCard(layout *logic.CardLayout, layoutName string, data *stats.ResponsePayload, printer func(string) string) *block.Block {
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
			ses, ok := data.Session.Rating.Ratings[wn8.WN8]
			if !ok {
				ses = -1
			}
			alt, ok := data.Snapshot.Rating.Ratings[wn8.WN8]
			if !ok {
				alt = -1
			}
			block := WN8BlockFromStats(layoutName, b, ses, alt, printer)
			if block != nil {
				content = append(content, *block)
			}

		default:
			block := BlockFromStats(layoutName, b, data.Session.Rating.Total, data.Snapshot.Rating.Total, printer)
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

func GenerateRandomOverviewCard(layout *logic.CardLayout, layoutName string, data *stats.ResponsePayload, printer func(string) string) *block.Block {
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
			ses, ok := data.Session.Regular.Ratings[wn8.WN8]
			if !ok {
				ses = -1
			}
			alt, ok := data.Snapshot.Regular.Ratings[wn8.WN8]
			if !ok {
				alt = -1
			}
			block := WN8BlockFromStats(layoutName, b, ses, alt, printer)
			if block != nil {
				content = append(content, *block)
			}

		default:
			block := BlockFromStats(layoutName, b, data.Session.Regular.Total, data.Snapshot.Regular.Total, printer)
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
