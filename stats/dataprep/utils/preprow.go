package utils

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/icons"
	dataprep "byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Formatted string for all time will fallback to session if empty
type FmtStr struct {
	Session string
	AllTime string
}

// Is _should_ be safe to pass nil values
func PrepContentRows(input dataprep.DataprepInput, fmtString FmtStr, isPercentage bool, sessionValue, sessionOf, allTimeValue, allTimeOf int, tags ...string) []types.StatsBlockRow {
	var rows []types.StatsBlockRow

	var sessionRow types.StatsBlockRow
	var allTimeRow types.StatsBlockRow

	var sessionString string = "-"
	var sessionFloat float64 = 0
	if sessionOf > 0 {
		sessionFloat = divideValueOf(sessionValue, sessionOf, isPercentage)
		sessionString = fmt.Sprintf(fmt.Sprintf(fmtString.Session, getFormattedString(isPercentage)), sessionFloat)
	}

	var allTimeString string = "-"
	var allTimeFloat float64 = 0
	if allTimeOf > 0 {
		allTimeFloat = divideValueOf(allTimeValue, allTimeOf, isPercentage)
		fmtStringFixed := fmtString.AllTime
		if fmtStringFixed == "" {
			fmtStringFixed = fmtString.Session
		}
		allTimeString = fmt.Sprintf(fmt.Sprintf(fmtStringFixed, getFormattedString(isPercentage)), allTimeFloat)
	}

	// Session icon
	if input.Options.WithIcons && input.Options.Block.HasIcon {
		var icon types.ContentIcon
		icon.Color = input.Options.Block.IconColorOverWrite
		iconsDict := input.Options.Block.IconDictOverwrite
		if iconsDict == nil {
			iconsDict = icons.IconsArrows
		}

		if sessionFloat > allTimeFloat {
			if icon.Color == "" {
				icon.Color = icons.IconColorGreen
				if sessionFloat/allTimeFloat > 1.6 {
					icon.Color = icons.IconColorPurple
				} else if sessionFloat/allTimeFloat > 1.4 {
					icon.Color = icons.IconColorTeal
				}
			}
			icon.Name = iconsDict[icons.IconDirectionUp]
		} else if sessionFloat < allTimeFloat {
			if icon.Color == "" {
				icon.Color = icons.IconColorRed
				if sessionFloat/allTimeFloat > 0.9 {
					icon.Color = icons.IconColorYellow
				}
			}
			icon.Name = iconsDict[icons.IconDirectionDown]
		} else if icon.Color == "" {
			icon.Color = icons.IconColorNeutral
			icon.Name = icons.IconsLines[icons.IconDirectionLeft] // same as right and will be horizontal

		}

		sessionRow.Content = append(sessionRow.Content, types.StatsBlockRowContent{
			Tags:        []string{input.Options.Block.GenerationTag, TagIcon, TagSession},
			Content:     icon,
			Type:        types.ContentTypeIcon,
			IsLocalized: false,
		})
	}

	// Session content
	sessionRow.Content = append(sessionRow.Content, types.StatsBlockRowContent{
		Tags:        append([]string{input.Options.Block.GenerationTag, TagSession}, tags...),
		Type:        types.ContentTypeText,
		Content:     sessionString,
		IsLocalized: false,
	})
	rows = append(rows, sessionRow)

	// All time content
	if input.Options.WithAllTime {
		allTimeRow.Content = append(allTimeRow.Content, types.StatsBlockRowContent{
			Tags:        append([]string{input.Options.Block.GenerationTag, TagAllTime}, tags...),
			Type:        types.ContentTypeText,
			Content:     allTimeString,
			IsLocalized: false,
		})
		rows = append(rows, allTimeRow)
	}

	// Localized label
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: input.Options.Block.LocalizationTag,
		})
		rows = append(rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{input.Options.Block.GenerationTag, TagLabel},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}
	return rows
}

func getFormattedString(isPercentage bool) string {
	if isPercentage {
		return "%.2f%%"
	}
	return "%.0f"
}

func divideValueOf(value, of int, isPercent bool) float64 {
	if isPercent {
		return ((float64(value) / float64(of)) * 100)
	}
	return float64(value) / float64(of)
}
