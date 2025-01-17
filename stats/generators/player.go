package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	str "github.com/byvko-dev/am-core/helpers/strings"
	"github.com/byvko-dev/am-types/api/stats/v1"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func GeneratePlayerCard(layout *logic.CardLayout, layoutName string, data *stats.ResponsePayload, printer func(string) string) *block.Block {
	var card block.Block
	card.Style = layout.CardStyle
	card.ContentType = block.ContentTypeBlocks

	var blocks []block.Block
	for _, definition := range layout.Blocks {
		switch definition.ValueKind {
		case logic.PlayerName:
			layout := layouts.LoadDefinition(layoutName, definition)
			values := make(logic.Values)
			values[logic.String] = str.Or(data.Account.Nickname, "Unknown Player")
			layout.Values = values
			b := layout.ToBlock(printer)
			if b != nil {
				blocks = append(blocks, *b)
			}

		case logic.PlayerClanTag:
			if data.Account.Clan.Tag == "" {
				continue
			}
			layout := layouts.LoadDefinition(layoutName, definition)
			values := make(logic.Values)
			values[logic.String] = fmt.Sprintf("[%v]", data.Account.Clan.Tag)
			layout.Values = values
			b := layout.ToBlock(printer)
			if b != nil {
				blocks = append(blocks, *b)
			}
		default:
			continue
		}
	}
	if len(blocks) == 0 {
		return nil
	}

	nameBlock := block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     blocks,
		Style:       layout.ContentStyle,
	}

	card.Content = []block.Block{nameBlock}
	return &card

}
