package utils

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/icons"
	dataprep "byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"github.com/byvko-dev/am-types/dataprep/v1/block"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Formatted string for all time will fallback to session if empty
type FmtStr struct {
	Session string
	AllTime string
}

func PrepContentRows(input dataprep.DataprepInput, fmtString FmtStr, isPercentage bool, sessionValue, sessionOf, allTimeValue, allTimeOf int, tags ...string) []block.Block {
	var rows []block.Block

	var sessionRowContent []block.Block
	var allTimeRowContent []block.Block

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

	// Session content
	sessionRowContent = append(sessionRowContent, block.Block{
		Tags:        append([]string{input.Options.Block.GenerationTag, TagSession}, tags...),
		ContentType: block.ContentTypeText,
		Content:     sessionString,
	})

	// Session icon
	if input.Options.WithIcons && input.Options.Block.HasIcon {
		var icon block.Block
		icon.Style.Color = input.Options.Block.IconColorOverWrite
		iconsDict := input.Options.Block.IconDictOverwrite
		if iconsDict == nil {
			iconsDict = icons.IconsArrows
		}
		// TagInvisible

		if sessionFloat > allTimeFloat {
			if icon.Style.Color == "" {
				icon.Style.Color = icons.IconColorGreen
				if sessionFloat/allTimeFloat > 1.6 {
					icon.Style.Color = icons.IconColorPurple
				} else if sessionFloat/allTimeFloat > 1.4 {
					icon.Style.Color = icons.IconColorTeal
				}
			}
			icon.Content = iconsDict[icons.IconDirectionUp]
		} else if sessionFloat < allTimeFloat {
			if icon.Style.Color == "" {
				icon.Style.Color = icons.IconColorRed
				if sessionFloat/allTimeFloat > 0.9 {
					icon.Style.Color = icons.IconColorYellow
				}
			}
			icon.Content = iconsDict[icons.IconDirectionDown]
		} else if icon.Style.Color == "" {
			icon.Style.Color = icons.IconColorNeutral
			icon.Content = icons.IconsLines[icons.IconDirectionLeft] // same as right and will be horizontal
		}

		iconContent := block.Block{
			Tags:        []string{input.Options.Block.GenerationTag, TagIcon, TagSession},
			ContentType: block.ContentTypeIcon,
			Content:     icon,
		}

		sessionRowContent = append([]block.Block{iconContent}, sessionRowContent...)

		iconInvisible := iconContent
		iconInvisible.Tags = append(iconInvisible.Tags, TagInvisible)
		sessionRowContent = append(sessionRowContent, iconInvisible)
	}
	rows = append(rows, block.Block{
		Tags:        append([]string{input.Options.Block.GenerationTag, TagSession}, tags...),
		ContentType: block.ContentTypeText,
		Content:     sessionRowContent,
	})

	// All time content
	if input.Options.WithAllTime {
		allTimeRowContent = append(allTimeRowContent, block.Block{
			Tags:        append([]string{input.Options.Block.GenerationTag, TagAllTime}, tags...),
			ContentType: block.ContentTypeText,
			Content:     allTimeString,
		})
		rows = append(rows, block.Block{
			Tags:        append([]string{input.Options.Block.GenerationTag, TagAllTime}, tags...),
			ContentType: block.ContentTypeText,
			Content:     allTimeRowContent,
		})
	}

	// Localized label
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: input.Options.Block.LocalizationTag,
		})
		rows = append(rows, block.Block{
			Content: []block.Block{{
				Tags:        []string{input.Options.Block.GenerationTag, TagLabel},
				ContentType: block.ContentTypeText,
				Content:     label,
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
