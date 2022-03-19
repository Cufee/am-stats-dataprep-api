package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-core/logs"
	api "github.com/byvko-dev/am-types/stats/v1"
)

func GeneratePlayerCard(stats *api.PlayerRawStats, options types.PlayerOptions) (types.StatsCard, error) {
	if !options.WithClanTag && !options.WithName && !options.WithPins {
		return types.StatsCard{}, fmt.Errorf("no options provided")
	}

	var contentElements []types.StatsBlockRowContent
	if options.WithName {
		// Player name
		contentElements = append(contentElements, types.StatsBlockRowContent{
			Type:    types.ContentTypeText,
			Content: stats.PlayerDetails.Name,
		})
	}

	// Clan tag if it exists
	if options.WithClanTag && stats.PlayerDetails.ClanTag != "" {
		contentElements = append(contentElements, types.StatsBlockRowContent{
			Type:    types.ContentTypeText,
			Content: fmt.Sprintf("[%s]", stats.PlayerDetails.ClanTag),
		})
	}

	if options.WithPins {
		logs.Warning("GeneratePlayerCard: pins not implemented yet")
	}

	// Assemble the content
	var card types.StatsCard
	card.Rows = append(card.Rows, types.StatsCardRow{
		Blocks: []types.StatsBlock{{
			Rows: []types.StatsBlockRow{
				{Content: contentElements},
			},
			Tags: []string{},
		}},
	})

	return card, nil
}
