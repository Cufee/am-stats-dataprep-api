package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func WN8RatingBlock(input DataprepInput, ratingSession, ratingAllTime, sessionColor, allTimeColor string) (types.StatsBlock, error) {
	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockWN8Rating)

	var ratingSessionRow types.StatsBlockRow
	if sessionColor != "" {
		var icon types.ContentIcon
		icon.Color = sessionColor
		icon.Name = "circle"

		ratingSessionRow.Content = append(ratingSessionRow.Content, types.StatsBlockRowContent{
			Tags:        []string{types.BlockWN8Rating, "ratingSessionIcon", tagSession},
			Content:     icon,
			Type:        types.ContentTypeIcon,
			IsLocalized: false,
		})
	}
	ratingSessionRow.Content = append(ratingSessionRow.Content, types.StatsBlockRowContent{
		Tags:        []string{types.BlockWN8Rating, "ratingSession", tagSession},
		Content:     fmt.Sprint(ratingSession),
		Type:        types.ContentTypeText,
		IsLocalized: false,
	})
	block.Rows = append(block.Rows, ratingSessionRow)

	var ratingAllTimeRow types.StatsBlockRow
	if input.Options.WithAllTime {
		if allTimeColor != "" {
			var icon types.ContentIcon
			icon.Color = allTimeColor
			icon.Name = "circle"

			ratingAllTimeRow.Content = append(ratingAllTimeRow.Content, types.StatsBlockRowContent{
				Tags:        []string{types.BlockWN8Rating, "ratingAllTimeIcon", tagAllTime},
				Content:     icon,
				Type:        types.ContentTypeIcon,
				IsLocalized: false,
			})
		}
		ratingAllTimeRow.Content = append(ratingAllTimeRow.Content, types.StatsBlockRowContent{
			Tags:        []string{types.BlockWN8Rating, "ratingAllTime", tagAllTime},
			Content:     fmt.Sprint(ratingAllTime),
			Type:        types.ContentTypeText,
			IsLocalized: false,
		})
		block.Rows = append(block.Rows, ratingAllTimeRow)
	}

	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_wn8_rating",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockWN8Rating, "ratingLabel", tagLabel},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}

// GetRatingColor - Rating color calculator
func GetRatingColor(r int) string {
	if r > 0 && r < 301 {
		return "rgba(255, 0, 0, 0.7)"
	}
	if r > 300 && r < 451 {
		return "rgba(251, 83, 83, 0.7)"
	}
	if r > 450 && r < 651 {
		return "rgba(255, 160, 49, 0.7)"
	}
	if r > 650 && r < 901 {
		return "rgba(255, 244, 65, 0.7)"
	}
	if r > 900 && r < 1201 {
		return "rgba(149, 245, 62, 0.7)"
	}
	if r > 1200 && r < 1601 {
		return "rgba(103, 190, 51, 0.7)"
	}
	if r > 1600 && r < 2001 {
		return "rgba(106, 236, 255, 0.7)"
	}
	if r > 2000 && r < 2451 {
		return "rgba(46, 174, 193, 0.7)"
	}
	if r > 2450 && r < 2901 {
		return "rgba(208, 108, 255, 0.7)"
	}
	if r > 2900 {
		return "rgba(142, 65, 177, 0.7)"
	}
	return "rgba(255, 0, 0, 0)"
}
