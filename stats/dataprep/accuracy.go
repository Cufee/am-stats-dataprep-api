package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func ShotAccuracyBlock(input DataprepInput) (types.StatsBlock, error) {
	if input.Stats.Session == nil {
		return types.StatsBlock{}, fmt.Errorf("session stats are nil")
	}

	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockAverageDamage)

	var sessionAccuracy string = "-"
	if input.Stats.Session.Shots > 0 {
		sessionAccuracyFloat := ((float64(input.Stats.Session.Hits) / float64(input.Stats.Session.Shots)) * 100)
		sessionAccuracy = fmt.Sprintf("%.2f", sessionAccuracyFloat) + "%"

	}
	var allTimeAccuracy string = "-"
	if input.Stats.AllTime != nil && input.Stats.AllTime.Shots > 0 {
		allTimeAccuracyFloat := ((float64(input.Stats.AllTime.Hits) / float64(input.Stats.AllTime.Shots)) * 100)
		allTimeAccuracy = fmt.Sprintf("%.2f", allTimeAccuracyFloat) + "%"
	}

	block.Rows = append(block.Rows, types.StatsBlockRow{
		Content: []types.StatsBlockRowContent{{
			Tags:        []string{types.BlockShotAccuracy, "shotAccuracySession", tagSession},
			Content:     sessionAccuracy,
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}},
	})
	if input.Options.WithAllTime {
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockShotAccuracy, "shotAccuracyAllTime", tagAllTime},
				Content:     allTimeAccuracy,
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}},
		})
	}
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_shot_accuracy",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockShotAccuracy, "shotAccuracyLabel", tagLabel},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}
