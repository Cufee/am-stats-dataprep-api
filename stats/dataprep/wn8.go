package dataprep

// func WN8RatingBlock(input types.DataprepInput, ratingSession, ratingAllTime int) (block.Block, error) {
// 	var b block.Block
// 	ratingColor, level := getRatingColorLevel(ratingSession)
// 	if input.Options.Block.IconColorOverWrite == (color.RGBA{}) {
// 		input.Options.Block.IconColorOverWrite = ratingColor
// 	}
// 	if input.Options.Block.IconDictOverwrite == nil {
// 		input.Options.Block.IconDictOverwrite = icons.IconsRating(level)
// 	}
// 	if ratingSession < 0 {
// 		input.Options.WithAllTime = false
// 		input.Options.Block.HasIcon = false
// 	}
// 	fixTag := fmt.Sprintf("fixIcon-%v", input.Options.Block.HasIcon && input.Options.Block.HasInvisibleIcon)
// 	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, ratingSession, 1, ratingAllTime, 1)
// 	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, input.Options.Block.GenerationTag+"Block", fixTag))
// 	b.ContentType = block.ContentTypeBlocks
// 	b.Tags = append(b.Tags, fixTag)
// 	return b, nil
// }

// // GetRatingColor - Rating color calculator
// func getRatingColorLevel(r int) (color.RGBA, int) {
// 	if r > 0 && r < 301 {
// 		return helpers.HexToColor("#fb5353"), 1
// 	}
// 	// Yellow
// 	if r > 300 && r < 451 {
// 		return helpers.HexToColor("#ffa02f"), 2
// 	}
// 	if r > 450 && r < 651 {
// 		return helpers.HexToColor("#ffa02f"), 3
// 	}
// 	// Green
// 	if r > 650 && r < 901 {
// 		return helpers.HexToColor("#67be35"), 1
// 	}
// 	if r > 900 && r < 1201 {
// 		return helpers.HexToColor("#67be35"), 2
// 	}
// 	// Teal
// 	if r > 1200 && r < 1601 {
// 		return helpers.HexToColor("#6ae6ff"), 1
// 	}
// 	if r > 1600 && r < 2001 {
// 		return helpers.HexToColor("#6ae6ff"), 2
// 	}
// 	if r > 2000 && r < 2451 {
// 		return helpers.HexToColor("#6ae6ff"), 3
// 	}
// 	// Purple
// 	if r > 2450 && r < 2901 {
// 		return helpers.HexToColor("#9d53cf"), 1
// 	}
// 	if r > 2900 && r < 4501 {
// 		return helpers.HexToColor("#9d53cf"), 2
// 	}
// 	if r > 4500 && r < 6001 {
// 		return helpers.HexToColor("#9d53cf"), 3
// 	}
// 	if r > 6000 {
// 		return helpers.HexToColor("#FFD700"), 1
// 	}
// 	return helpers.HexToColor("#647893"), 1

// }
