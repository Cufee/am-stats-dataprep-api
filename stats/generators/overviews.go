package generators

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/logs"
	api "byvko.dev/repo/am-stats-dataprep-api/stats-api/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateOverviewCard(stats *api.PlayerRawStats, options types.OverviewOptions, localizer *i18n.Localizer) (types.StatsCard, error) {
	switch options.Type {
	case types.OverviewTypeRating:
		return generateRatingOverviewCard(stats, options, localizer)
	case types.OverviewTypeRegular:
		return generateRandomOverviewCard(stats, options, localizer)
	default:
		return types.StatsCard{}, fmt.Errorf("unknown overview type: %s", options.Type)
	}
}

func generateRatingOverviewCard(stats *api.PlayerRawStats, options types.OverviewOptions, localizer *i18n.Localizer) (types.StatsCard, error) {

	var row types.StatsCardRow
	for _, block := range options.Blocks {
		if block == types.BlockWN8Rating {
			var input dataprep.DataprepInput
			input.Options.WithAllTime = options.WithAllTimeStats
			input.Options.WithLabel = options.WithLabels
			input.Localizer = localizer

			sessionAccuracy := "-"
			if stats.SessionStats.StatsRating.Hits > 0 {
				sessionAccuracy = fmt.Sprintf("%.2f", float64(stats.SessionStats.StatsRating.Hits)/float64(stats.SessionStats.StatsRating.Shots)*100)
			}

			allTimeAccuracy := "-"
			if stats.PlayerDetails.Stats.Rating.Hits > 0 {
				allTimeAccuracy = fmt.Sprintf("%.2f", float64(stats.PlayerDetails.Stats.Rating.Hits)/float64(stats.PlayerDetails.Stats.Rating.Shots)*100)
			}

			block, err := dataprep.WN8RatingBlock(input, sessionAccuracy, allTimeAccuracy, "", "") // Needs to be replaced with accuracy as there is no WN8 in rating
			if err != nil {
				logs.Warning("generateRatingOverviewCard: error generating rating block for %v: %s", stats.PlayerDetails.ID, err)
				continue
			}
			row.Blocks = append(row.Blocks, block)
			continue
		}

		var input dataprep.DataprepInput
		input.Stats.Session = &stats.SessionStats.StatsRating
		input.Stats.AllTime = &stats.PlayerDetails.Stats.Rating
		input.Options.WithAllTime = options.WithAllTimeStats
		input.Options.WithLabel = options.WithLabels
		input.Localizer = localizer

		block, err := dataprep.BlockFromStats(input, block)
		if err != nil {
			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", stats.PlayerDetails.ID, err)
			continue
		}
		row.Blocks = append(row.Blocks, block)
	}

	var card types.StatsCard
	if options.WithTitle {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_rating_overview_title",
		})

		card.Rows = append(card.Rows, types.StatsCardRow{
			Blocks: []types.StatsBlock{
				{
					Rows: []types.StatsBlockRow{
						{
							Content: []types.StatsBlockRowContent{{
								Content: label,
								Tags:    []string{"overview_title"},
							}},
						},
					},
				},
			},
		})
	}
	card.Rows = append(card.Rows, row)
	return card, nil
}

func generateRandomOverviewCard(stats *api.PlayerRawStats, options types.OverviewOptions, localizer *i18n.Localizer) (types.StatsCard, error) {
	var row types.StatsCardRow
	for _, block := range options.Blocks {
		if block == types.BlockWN8Rating {
			var input dataprep.DataprepInput
			input.Options.WithAllTime = options.WithAllTimeStats
			input.Options.WithLabel = options.WithLabels
			input.Localizer = localizer

			sessionRating := "-"
			if stats.SessionStats.SessionRating > 0 {
				sessionRating = fmt.Sprintf("%d", stats.SessionStats.SessionRating)
			}
			allTimeRating := "-"
			if stats.PlayerCache.CareerWN8 > 0 {
				allTimeRating = fmt.Sprintf("%d", stats.PlayerCache.CareerWN8)
			}

			block, err := dataprep.WN8RatingBlock(input, sessionRating, allTimeRating, dataprep.GetRatingColor(stats.SessionStats.SessionRating), dataprep.GetRatingColor(stats.PlayerCache.CareerWN8))
			if err != nil {
				logs.Warning("generateRatingOverviewCard: error generating rating block for %v: %s", stats.PlayerDetails.ID, err)
				continue
			}
			row.Blocks = append(row.Blocks, block)
			continue
		}

		var input dataprep.DataprepInput
		input.Stats.Session = &stats.SessionStats.StatsAll
		input.Stats.AllTime = &stats.PlayerDetails.Stats.All
		input.Options.WithAllTime = options.WithAllTimeStats
		input.Options.WithLabel = options.WithLabels
		input.Localizer = localizer

		block, err := dataprep.BlockFromStats(input, block)
		if err != nil {
			logs.Warning("generateRatingOverviewCard: error generating block for %v: %s", stats.PlayerDetails.ID, err)
			continue
		}
		row.Blocks = append(row.Blocks, block)
	}

	var card types.StatsCard
	if options.WithTitle {
		label, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_random_overview_title",
		})

		card.Rows = append(card.Rows, types.StatsCardRow{
			Blocks: []types.StatsBlock{
				{
					Rows: []types.StatsBlockRow{
						{
							Content: []types.StatsBlockRowContent{{
								Content: label,
								Tags:    []string{"overview_title"},
							}},
						},
					},
				},
			},
		})
	}
	card.Rows = append(card.Rows, row)
	return card, nil
}
