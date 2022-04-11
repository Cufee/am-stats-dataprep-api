package generators

// func GeneratePlayerCard(stats *api.PlayerRawStats, options settings.PlayerOptions, styleName string) (block.Block, error) {
// 	if !options.WithClanTag && !options.WithName && !options.WithPins {
// 		return block.Block{}, fmt.Errorf("no options provided")
// 	}

// 	var contentElements []block.Block
// 	if options.WithName {
// 		// Player name
// 		contentElements = append(contentElements, block.Block{
// 			Style:       styles.LoadWithTags(styleName, utils.TagPlayerName),
// 			ContentType: block.ContentTypeText,
// 			Content:     stats.PlayerDetails.Name,
// 		})
// 	}

// 	// Clan tag if it exists
// 	if options.WithClanTag && stats.PlayerDetails.ClanTag != "" {
// 		contentElements = append(contentElements, block.Block{
// 			Style:       styles.LoadWithTags(styleName, utils.TagPlayerClan),
// 			ContentType: block.ContentTypeText,
// 			Content:     fmt.Sprintf("[%s]", stats.PlayerDetails.ClanTag),
// 		})
// 	}

// 	if options.WithPins {
// 		logs.Warning("GeneratePlayerCard: pins not implemented yet")
// 	}

// 	fullName := []block.Block{{
// 		ContentType: block.ContentTypeBlocks,
// 		Content:     contentElements,
// 		Style:       styles.LoadWithTags(styleName, "playerName"),
// 	}}

// 	return block.Block{
// 		Style:       styles.LoadWithTags(styleName, "card"),
// 		ContentType: block.ContentTypeBlocks,
// 		Content: []block.Block{{
// 			ContentType: block.ContentTypeBlocks,
// 			Content:     fullName,
// 			Style:       shared.AlignVertical.Merge(styles.LoadWithTags(styleName, "playerNameContainer")),
// 		}}}, nil
// }
