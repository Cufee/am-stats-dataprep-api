package dataprep

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func BattlesBlock(input DataprepInput) (types.StatsBlock, error) {
	if input.Stats.Session == nil {
		return types.StatsBlock{}, fmt.Errorf("session stats are nil")
	}

	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockBattles)

	var battlesSession string = "-"
	if input.Stats.Session.Battles > 0 {
		battlesSession = fmt.Sprintf("%d", input.Stats.Session.Battles)
	}
	block.Rows = append(block.Rows, types.StatsBlockRow{
		Content: []types.StatsBlockRowContent{{
			Tags:        []string{types.BlockBattles, "battlesSession", tagSession},
			Content:     battlesSession,
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}},
	})

	if input.Options.WithAllTime {
		var battlesAllTime string = "-"
		if input.Stats.AllTime != nil && input.Stats.AllTime.Battles > 0 {
			battlesAllTime = fmt.Sprintf("%d", input.Stats.AllTime.Battles)
		}
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockBattles, "battlesAllTime", tagAllTime},
				Content:     battlesAllTime,
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}},
		})
	}
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_battles",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockBattles, "battlesLabel", tagLabel},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}

func WinrateBlock(input DataprepInput) (types.StatsBlock, error) {
	if input.Stats.Session == nil {
		return types.StatsBlock{}, fmt.Errorf("session stats are nil")
	}

	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockWinrate)

	var winrateSession string = "-"
	if input.Stats.Session.Battles > 0 {
		if input.Stats.Session.Wins > 0 {
			winrateSessionFloat := ((float64(input.Stats.Session.Wins) / float64(input.Stats.Session.Battles)) * 100)
			winrateSession = fmt.Sprintf("%.2f", winrateSessionFloat) + "%"
		} else {
			winrateSession = "0%"
		}
	}

	var winrateAllTime string = "-"
	if input.Stats.AllTime != nil && input.Stats.AllTime.Battles > 0 {
		if input.Stats.AllTime.Wins > 0 {
			winrateAllTimeFloat := ((float64(input.Stats.AllTime.Wins) / float64(input.Stats.AllTime.Battles)) * 100)
			winrateAllTime = fmt.Sprintf("%.2f", winrateAllTimeFloat) + "%"
		} else {
			winrateAllTime = "0%"
		}
	}
	block.Rows = append(block.Rows, types.StatsBlockRow{
		Content: []types.StatsBlockRowContent{{
			Tags:        []string{types.BlockWinrate, "winrateSession", tagSession},
			Content:     winrateSession,
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}},
	})

	if input.Options.WithAllTime {
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockWinrate, "winrateAllTime", tagAllTime},
				Content:     winrateAllTime,
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}},
		})
	}
	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_winrate",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockWinrate, "winrateLabel", "label"},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}

func WinrateWithBattlesBlock(input DataprepInput) (types.StatsBlock, error) {
	if input.Stats.Session == nil {
		return types.StatsBlock{}, fmt.Errorf("session stats are nil")
	}

	var block types.StatsBlock
	block.Tags = append(block.Tags, "block"+types.BlockWinrateWithBattles)

	var winrateSession string = "-"
	if input.Stats.Session.Battles > 0 {
		if input.Stats.Session.Wins > 0 {
			winrateSessionFloat := ((float64(input.Stats.Session.Wins) / float64(input.Stats.Session.Battles)) * 100)
			winrateSession = fmt.Sprintf("%.2f", winrateSessionFloat) + "%"
		} else {
			winrateSession = "0%"
		}
	}
	block.Rows = append(block.Rows, types.StatsBlockRow{
		Content: []types.StatsBlockRowContent{{
			Tags:        []string{types.BlockWinrateWithBattles, "winrateWithBattlesSession", types.BlockWinrateWithBattles + "_winrate", tagSession},
			Content:     winrateSession,
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}, {
			Tags:        []string{types.BlockWinrateWithBattles, "winrateWithBattlesSessionBattles", types.BlockWinrateWithBattles + "_battles", tagSession},
			Content:     fmt.Sprintf("(%d)", input.Stats.Session.Battles),
			Type:        types.ContentTypeText,
			IsLocalized: false,
		}},
	})

	var winrateAllTime string = "-"
	if input.Stats.AllTime != nil && input.Stats.AllTime.Battles > 0 {
		if input.Stats.AllTime.Wins > 0 {
			winrateAllTimeFloat := ((float64(input.Stats.AllTime.Wins) / float64(input.Stats.AllTime.Battles)) * 100)
			winrateAllTime = fmt.Sprintf("%.2f", winrateAllTimeFloat) + "%"
		} else {
			winrateAllTime = "0%"
		}
	}
	if input.Options.WithAllTime {
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockWinrateWithBattles, "winrateWithBattlesAllTime", types.BlockWinrateWithBattles + "_winrate", tagAllTime},
				Content:     winrateAllTime,
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}, {
				Tags:        []string{types.BlockWinrateWithBattles, "winrateWithBattlesAllTime", types.BlockWinrateWithBattles + "_battles", tagAllTime},
				Content:     fmt.Sprintf("(%d)", input.Stats.AllTime.Battles),
				Type:        types.ContentTypeText,
				IsLocalized: false,
			}},
		})
	}

	if input.Options.WithLabel {
		label, _ := input.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "localized_winrate_with_battles",
		})
		block.Rows = append(block.Rows, types.StatsBlockRow{
			Content: []types.StatsBlockRowContent{{
				Tags:        []string{types.BlockWinrateWithBattles, "winrateWithBattlesLabel", "label"},
				Content:     label,
				Type:        types.ContentTypeText,
				IsLocalized: true,
			}},
		})
	}

	return block, nil
}
