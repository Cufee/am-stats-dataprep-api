package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	stats "byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/dataprep/v1/block"
	"github.com/byvko-dev/am-types/dataprep/v1/settings"
	api "github.com/byvko-dev/am-types/stats/v1"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer) (block.Block, error) {
	switch options.Type {
	case settings.OverviewTypeRating:
		return generateRatingOverviewCard(statsData, options, localizer)
	case settings.OverviewTypeRegular:
		return generateRandomOverviewCard(statsData, options, localizer)
	default:
		return block.Block{}, fmt.Errorf("unknown overview type: %s", options.Type)
	}
}

func generateRatingOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer) (block.Block, error) {
	var rowContent []block.Block
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
		rowContent = append(rowContent, block)
	}

	var cardRows []block.Block
	if options.WithTitle {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_rating_overview_title",
		})

		cardRows = append(cardRows, block.Block{
			ContentType: block.ContentTypeBlocks,
			Content: []block.Block{
				{
					ContentType: block.ContentTypeBlocks,
					Content: []block.Block{
						{
							ContentType: block.ContentTypeBlocks,
							Content: []block.Block{{
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
	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
	})
	return block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     cardRows,
	}, nil
}

func generateRandomOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer) (block.Block, error) {
	var rowContent []block.Block
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
			rowContent = append(rowContent, block)
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
		rowContent = append(rowContent, block)
	}

	var cardRows []block.Block
	if options.WithTitle {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_random_overview_title",
		})

		cardRows = append(cardRows, block.Block{
			ContentType: block.ContentTypeBlocks,
			Content: []block.Block{
				{
					ContentType: block.ContentTypeBlocks,
					Content: []block.Block{
						{
							ContentType: block.ContentTypeBlocks,
							Content: []block.Block{{
								ContentType: block.ContentTypeText,
								Content:     label,
								Tags:        []string{"overview_title"},
							}},
						},
					},
				},
			},
			Tags: []string{"overview_title_row"},
		})
	}
	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
	})
	return block.Block{
		ContentType: block.ContentTypeBlocks,
		Style: block.Style{
			AlignItems: block.AlignItemsVertical,
		},
		Content: cardRows,
	}, nil
}
