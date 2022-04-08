package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"
	stats "byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
	api "github.com/byvko-dev/am-types/stats/v1"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer, styleName string) (block.Block, error) {
	switch options.Type {
	case settings.OverviewTypeRating:
		return generateRatingOverviewCard(statsData, options, localizer, styleName)
	case settings.OverviewTypeRegular:
		return generateRandomOverviewCard(statsData, options, localizer, styleName)
	default:
		return block.Block{}, fmt.Errorf("unknown overview type: %s", options.Type)
	}
}

func generateRatingOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer, styleName string) (block.Block, error) {
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
		input.Options.Style = styleName
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
			Content: []block.Block{{
				ContentType: block.ContentTypeText,
				Content:     label,
				Tags:        []string{"overview_title"},
				Style:       styles.LoadWithTags(styleName, "overview_title"),
			}},
			Tags:  []string{"overview_title_row", "overview_title"},
			Style: styles.LoadWithTags(styleName, "overview_title_row", "overview_title"),
		})
	}
	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
		Tags:        []string{"rating_overview_rows", "growX"},
		Style:       styles.LoadWithTags(styleName, "growX", "rating_overview_rows"),
	})
	cardContent := block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     cardRows,
		Tags:        []string{"rating_overview", "growX", "gap50"},
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "growX", "rating_overview", "gap50")),
	}

	return block.Block{
		Tags:        []string{"card"},
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "card")),
		ContentType: block.ContentTypeBlocks,
		Content:     []block.Block{cardContent},
	}, nil
}

func generateRandomOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer, styleName string) (block.Block, error) {
	var rowContent []block.Block
	for _, block := range options.Blocks {
		if block.GenerationTag == stats.BlockWN8Rating.GenerationTag {
			if statsData.SessionStats.SessionRating < 0 && statsData.PlayerDetails.CareerWN8 < 0 {
				continue
			}
			var input types.DataprepInput
			input.Options.WithAllTime = options.WithAllTimeStats
			input.Options.WithIcons = options.WithIcons
			input.Options.WithLabel = options.WithLabels
			input.Options.Style = styleName
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
		input.Options.Style = styleName
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
			ContentType: block.ContentTypeText,
			Content:     label,
			Tags:        []string{"overview_title"},
			Style:       styles.LoadWithTags(styleName, "overview_title_row"),
		})
	}
	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
		Tags:        []string{"overview_stats_row", "growX", "gap50"},
		Style:       styles.LoadWithTags(styleName, "overview_stats_row", "growX", "gap50"),
	})
	cardContent := block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     cardRows,
		Tags:        []string{"random_overview", "growX"},
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "random_overview", "growX")),
	}
	return block.Block{
		Tags:        []string{"card"},
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "card")),
		ContentType: block.ContentTypeBlocks,
		Content:     []block.Block{cardContent},
	}, nil
}
