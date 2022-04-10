package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	tags "byvko.dev/repo/am-stats-dataprep-api/stats/presets/shared"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"
	"github.com/byvko-dev/am-core/helpers/slices"
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
		if block.GenerationTag == tags.GenerationTagWN8Rating {
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
				Style:       styles.LoadWithTags(styleName, "overviewTitle"),
			}},
			Style: styles.LoadWithTags(styleName, "overviewTitleRow"),
		})
	}

	// Find fixtag
	fixTag := "fixIcon-false"
	for _, b := range rowContent {
		if slices.Contains(b.Tags, "fixIcon-true") > -1 {
			fixTag = "fixIcon-true"
			break
		}
	}

	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
		Style:       styles.LoadWithTags(styleName, "growX", "ratingOverviewRows", "statsContent", fixTag),
	})
	cardContent := block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     cardRows,
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "growX", "ratingOverview", "gap50", "statsContent")),
	}

	return block.Block{
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "card")),
		ContentType: block.ContentTypeBlocks,
		Content:     []block.Block{cardContent},
	}, nil
}

func generateRandomOverviewCard(statsData *api.PlayerRawStats, options settings.OverviewOptions, localizer *i18n.Localizer, styleName string) (block.Block, error) {
	var rowContent []block.Block
	for _, block := range options.Blocks {
		if block.GenerationTag == tags.GenerationTagWN8Rating {
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
			Style:       styles.LoadWithTags(styleName, "overviewTitle"),
		})
	}

	// Find fixtag
	fixTag := "fixIcon-false"
	for _, b := range rowContent {
		if slices.Contains(b.Tags, "fixIcon-true") > -1 {
			fixTag = "fixIcon-true"
			break
		}
	}
	cardRows = append(cardRows, block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     rowContent,
		Style:       styles.LoadWithTags(styleName, "overviewStatsRow", "statsContent", fixTag),
	})
	cardContent := block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     cardRows,
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "randomOverview")),
	}
	return block.Block{
		Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "card")),
		ContentType: block.ContentTypeBlocks,
		Content:     []block.Block{cardContent},
	}, nil
}
