package generators

import (
	"fmt"

	api "byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func GeneratePlayerCard(stats api.PlayerRawStats, options types.PlayerOptions) (types.StatsCard, error) {
	// Player name
	var contentElements []types.StatsBlockRowContent
	contentElements = append(contentElements, types.StatsBlockRowContent{
		Type:    types.ContentTypeText,
		Content: stats.PlayerDetails.Name,
	})

	// Clan tag if it exists
	if stats.PlayerDetails.ClanTag != "" {
		contentElements = append(contentElements, types.StatsBlockRowContent{
			Type:    types.ContentTypeText,
			Content: fmt.Sprintf("[%s]", stats.PlayerDetails.ClanTag),
		})
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
