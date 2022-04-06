package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
	api "github.com/byvko-dev/am-types/stats/v1"
)

func GeneratePlayerCard(stats *api.PlayerRawStats, options settings.PlayerOptions, styleName string) (block.Block, error) {
	if !options.WithClanTag && !options.WithName && !options.WithPins {
		return block.Block{}, fmt.Errorf("no options provided")
	}

	var contentElements []block.Block
	if options.WithName {
		// Player name
		contentElements = append(contentElements, block.Block{
			Tags:        []string{utils.TagPlayerName},
			Style:       styles.LoadWithTags(styleName, utils.TagPlayerName),
			ContentType: block.ContentTypeText,
			Content:     stats.PlayerDetails.Name,
		})
	}

	// Clan tag if it exists
	if options.WithClanTag && stats.PlayerDetails.ClanTag != "" {
		contentElements = append(contentElements, block.Block{
			Tags:        []string{utils.TagPlayerClan},
			Style:       styles.LoadWithTags(styleName, utils.TagPlayerClan),
			ContentType: block.ContentTypeText,
			Content:     fmt.Sprintf("[%s]", stats.PlayerDetails.ClanTag),
		})
	}

	if options.WithPins {
		logs.Warning("GeneratePlayerCard: pins not implemented yet")
	}

	return block.Block{
		Tags:        []string{"card"},
		Style:       styles.LoadWithTags(styleName, "card"),
		ContentType: block.ContentTypeBlocks,
		Content: []block.Block{{
			ContentType: block.ContentTypeBlocks,
			Content:     contentElements,
			Style:       styles.LoadWithTags(styleName, "player_name"),
			Tags:        []string{"player_name"},
		}}}, nil
}
