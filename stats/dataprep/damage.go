package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func AvarageDamageBlock(input DataprepInput) (types.StatsBlock, error) {
	if input.Stats.Session == nil {
		return types.StatsBlock{}, fmt.Errorf("session stats are nil")
	}

	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockAverageDamage)

	var averageDamageSession string = "-"
	if input.Stats.Session.Battles > 0 {
		averageDamageSession = fmt.Sprintf("%d", input.Stats.Session.DamageDealt/input.Stats.Session.Battles)
	}
	var averageDamageAllTime string = "-"
	if input.Stats.AllTime != nil && input.Stats.AllTime.Battles > 0 {
		averageDamageAllTime = fmt.Sprintf("%d", input.Stats.AllTime.DamageDealt/input.Stats.AllTime.Battles)
	}

	block.Rows = append(block.Rows, types.StatsBlockRow{
		Content: []types.StatsBlockRowContent{{
			Tags:        []string{types.BlockAverageDamage, "averageDamageSession", tagSession},
			Content:     averageDamageSession,
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}},
	})
	if input.Options.WithAllTime {
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockAverageDamage, "averageDamageAllTime", tagAllTime},
				Content:     averageDamageAllTime,
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}},
		})
	}
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_average_damage",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockAverageDamage, "averageDamageLabel", tagLabel},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}

func DamageDoneBlock(input DataprepInput) (types.StatsBlock, error) {
	if input.Stats.Session == nil {
		return types.StatsBlock{}, fmt.Errorf("session stats are nil")
	}

	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockDamageDone)

	var damageDoneSession string = "-"
	if input.Stats.Session.Battles > 0 {
		damageDoneSession = fmt.Sprintf("%d", input.Stats.Session.DamageDealt)
	}
	var damageDoneAllTime string = "-"
	if input.Stats.AllTime != nil && input.Stats.AllTime.Battles > 0 {
		damageDoneAllTime = fmt.Sprintf("%d", input.Stats.AllTime.DamageDealt)
	}

	block.Rows = append(block.Rows, types.StatsBlockRow{
		Content: []types.StatsBlockRowContent{{
			Tags:        []string{types.BlockDamageDone, "damageDoneSession", tagSession},
			Content:     damageDoneSession,
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}},
	})
	if input.Options.WithAllTime {
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockDamageDone, "damageDoneAllTime", tagAllTime},
				Content:     damageDoneAllTime,
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}},
		})
	}
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_damage_done",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockDamageDone, "damageDoneLabel", tagLabel},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}
