package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	stats "byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-core/logs"
	api "github.com/byvko-dev/am-types/stats/v1"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateOverviewCard(statsData *api.PlayerRawStats, options stats.OverviewOptions, localizer *i18n.Localizer) (stats.StatsCard, error) {
	switch options.Type {
	case stats.OverviewTypeRating:
		return generateRatingOverviewCard(statsData, options, localizer)
	case stats.OverviewTypeRegular:
		return generateRandomOverviewCard(statsData, options, localizer)
	default:
		return stats.StatsCard{}, fmt.Errorf("unknown overview type: %s", options.Type)
	}
}

func generateRatingOverviewCard(statsData *api.PlayerRawStats, options stats.OverviewOptions, localizer *i18n.Localizer) (stats.StatsCard, error) {
	var row stats.StatsCardRow
	for _, block := range options.Blocks {
		if block.GenerationTag == stats.BlockWN8Rating.GenerationTag {
			logs.Warning("generateRatingOverviewCard: error generating rating block for %v: rating battles have no WN8", statsData.PlayerDetails.ID)
			continue
		}

		var input types.DataprepInput
		input.Stats.Session = statsData.SessionStats.StatsRating
		input.Stats.AllTime = statsData.PlayerDetails.Stats.Rating
		input.Options.WithAllTime = options.WithAllTimeStats
		input.Options.WithLabel = options.WithLabels
		input.Options.WithIcons = options.WithIcons
		input.Options.Block = block
		input.Localizer = localizer

		block, err := dataprep.BlockFromStats(input)
		if err != nil {
			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", statsData.PlayerDetails.ID, err)
			continue
		}
		row.Blocks = append(row.Blocks, block)
	}

	var card stats.StatsCard
	if options.WithTitle {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_rating_overview_title",
		})

		card.Rows = append(card.Rows, stats.StatsCardRow{
			Blocks: []stats.StatsBlock{
				{
					Rows: []stats.StatsBlockRow{
						{
							Content: []stats.StatsBlockRowContent{{
								Content: label,
								Tags:    []string{"overview_title"},
							}},
						},
					},
				},
			},
			Tags: []string{"overview_title_row"},
		})
	}
	card.Rows = append(card.Rows, row)
	return card, nil
}

func generateRandomOverviewCard(statsData *api.PlayerRawStats, options stats.OverviewOptions, localizer *i18n.Localizer) (stats.StatsCard, error) {
	var row stats.StatsCardRow
	for _, block := range options.Blocks {
		if block.GenerationTag == stats.BlockWN8Rating.GenerationTag {
			var input types.DataprepInput
			input.Options.WithAllTime = options.WithAllTimeStats
			input.Options.WithIcons = options.WithIcons
			input.Options.WithLabel = options.WithLabels
			input.Options.Block = block
			input.Localizer = localizer
			block, err := dataprep.WN8RatingBlock(input, statsData.SessionStats.SessionRating, statsData.PlayerDetails.CareerWN8)
			if err != nil {
				logs.Warning("generateRatingOverviewCard: error generating rating block for %v: %s", statsData.PlayerDetails.ID, err)
				continue
			}
			row.Blocks = append(row.Blocks, block)
			continue
		}

		var input types.DataprepInput
		input.Stats.Session = statsData.SessionStats.StatsAll
		input.Stats.AllTime = statsData.PlayerDetails.Stats.All
		input.Options.WithAllTime = options.WithAllTimeStats
		input.Options.WithLabel = options.WithLabels
		input.Options.WithIcons = options.WithIcons
		input.Options.Block = block
		input.Localizer = localizer

		block, err := dataprep.BlockFromStats(input)
		if err != nil {
			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", statsData.PlayerDetails.ID, err)
			continue
		}
		row.Blocks = append(row.Blocks, block)
	}

	var card stats.StatsCard
	if options.WithTitle {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_random_overview_title",
		})

		card.Rows = append(card.Rows, stats.StatsCardRow{
			Blocks: []stats.StatsBlock{
				{
					Rows: []stats.StatsBlockRow{
						{
							Content: []stats.StatsBlockRowContent{{
								Content: label,
								Tags:    []string{"overview_title"},
							}},
						},
					},
				},
			},
			Tags: []string{"overview_title_row"},
		})
	}
	card.Rows = append(card.Rows, row)
	return card, nil
}
