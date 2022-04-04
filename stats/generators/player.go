package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/dataprep/v1/block"
	"github.com/byvko-dev/am-types/dataprep/v1/settings"
	api "github.com/byvko-dev/am-types/stats/v1"
)

func GeneratePlayerCard(stats *api.PlayerRawStats, options settings.PlayerOptions) (block.Block, error) {
	if !options.WithClanTag && !options.WithName && !options.WithPins {
		return block.Block{}, fmt.Errorf("no options provided")
	}

	var contentElements []block.Block
	if options.WithName {
		// Player name
		contentElements = append(contentElements, block.Block{
			Tags:        []string{utils.TagPlayerName},
			ContentType: block.ContentTypeText,
			Content:     stats.PlayerDetails.Name,
		})
	}

	// Clan tag if it exists
	if options.WithClanTag && stats.PlayerDetails.ClanTag != "" {
		contentElements = append(contentElements, block.Block{
			Tags:        []string{utils.TagPlayerClan},
			ContentType: block.ContentTypeText,
			Content:     fmt.Sprintf("[%s]", stats.PlayerDetails.ClanTag),
		})
	}

	if options.WithPins {
		logs.Warning("GeneratePlayerCard: pins not implemented yet")
	}

	// Assemble the content
	var card block.Block
	card.ContentType = block.ContentTypeBlocks
	card.Content = []block.Block{
		{
			ContentType: block.ContentTypeBlocks,
			Content: []block.Block{{
				ContentType: block.ContentTypeBlocks,
				Content: []block.Block{
					{
						ContentType: block.ContentTypeBlocks,
						Content:     contentElements,
					},
				},
			}},
		},
	}

	return card, nil
}
