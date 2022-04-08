package dataprep

import (
	"image/color"

	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"byvko.dev/repo/am-stats-dataprep-api/stats/helpers"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles"
	"byvko.dev/repo/am-stats-dataprep-api/stats/styles/shared"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
)

func WN8RatingBlock(input types.DataprepInput, ratingSession, ratingAllTime int) (block.Block, error) {
	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	input.Options.Block.IconColorOverWrite = getRatingColor(ratingSession)
	if ratingSession < 0 {
		input.Options.Block.HasIcon = false
	}
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, ratingSession, 1, ratingAllTime, 1)
	b.Style = shared.AlignVertical.Merge(styles.LoadWithTags(input.Options.Style, b.Tags...))
	b.ContentType = block.ContentTypeBlocks
	return b, nil
}

// GetRatingColor - Rating color calculator
func getRatingColor(r int) color.RGBA {
	if r > 0 && r < 301 {
		// return "rgba(255, 0, 0, 0.7)"
		// return color.RGBA{R: 255, G: 0, B: 0, A: 0x7f}
		return helpers.HexToColor("#ff0000")
	}
	if r > 300 && r < 451 {
		// return "rgba(251, 83, 83, 0.7)"
		// return color.RGBA{R: 251, G: 83, B: 83, A: 0x7f}
		return helpers.HexToColor("#fb5353")
	}
	if r > 450 && r < 651 {
		// return "rgba(255, 160, 49, 0.7)"
		// return color.RGBA{R: 255, G: 160, B: 49, A: 0x7f}
		return helpers.HexToColor("#ffa02f")
	}
	if r > 650 && r < 901 {
		// return "rgba(255, 244, 65, 0.7)"
		// return color.RGBA{R: 255, G: 244, B: 65, A: 0x7f}
		return helpers.HexToColor("#fff44d")
	}
	if r > 900 && r < 1201 {
		// return "rgba(149, 245, 62, 0.7)"
		// return color.RGBA{R: 149, G: 245, B: 62, A: 0x7f}
		return helpers.HexToColor("#95f442")
	}
	if r > 1200 && r < 1601 {
		// return "rgba(103, 190, 51, 0.7)"
		// return color.RGBA{R: 103, G: 190, B: 51, A: 0x7f}
		return helpers.HexToColor("#67be35")
	}
	if r > 1600 && r < 2001 {
		// return "rgba(106, 236, 255, 0.7)"
		// return color.RGBA{R: 106, G: 236, B: 255, A: 0x7f}
		return helpers.HexToColor("#6ae6ff")
	}
	if r > 2000 && r < 2451 {
		// return "rgba(46, 174, 193, 0.7)"
		// return color.RGBA{R: 46, G: 174, B: 193, A: 0x7f}
		return helpers.HexToColor("#2eafc1")
	}
	if r > 2450 && r < 2901 {
		// return "rgba(208, 108, 255, 0.7)"
		// return color.RGBA{R: 208, G: 108, B: 255, A: 0x7f}
		return helpers.HexToColor("#d062ff")
	}
	if r > 2900 {
		// return "rgba(142, 65, 177, 0.7)"
		// return color.RGBA{R: 142, G: 65, B: 177, A: 0x7f}
		return helpers.HexToColor("#8e41b1")
	}
	// return color.RGBA{R: 100, G: 116, B: 139, A: 0x7f}
	return helpers.HexToColor("#647893")

}
