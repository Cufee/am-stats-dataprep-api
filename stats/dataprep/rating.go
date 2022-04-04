package dataprep

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/utils"
	"github.com/byvko-dev/am-types/dataprep/v1/block"
)

func WN8RatingBlock(input types.DataprepInput, ratingSession, ratingAllTime int) (block.Block, error) {
	var b block.Block
	b.Tags = append(b.Tags, input.Options.Block.GenerationTag+"Block")
	input.Options.Block.IconColorOverWrite = getRatingColor(ratingSession)
	b.Content = utils.PrepContentRows(input, utils.FmtStr{Session: "%v"}, false, ratingSession, 1, ratingAllTime, 1)
	b.Style.AlignItems = block.AlignItemsVertical
	b.ContentType = block.ContentTypeBlocks
	return b, nil
}

// GetRatingColor - Rating color calculator
func getRatingColor(r int) string {
	if r > 0 && r < 301 {
		// return "rgba(255, 0, 0, 0.7)"
		return "#64748b"
	}
	if r > 300 && r < 451 {
		// return "rgba(251, 83, 83, 0.7)"
		return "#dc2626"
	}
	if r > 450 && r < 651 {
		// return "rgba(255, 160, 49, 0.7)"
		return "#ef4444"
	}
	if r > 650 && r < 901 {
		// return "rgba(255, 244, 65, 0.7)"
		return "#fb923c"
	}
	if r > 900 && r < 1201 {
		// return "rgba(149, 245, 62, 0.7)"
		return "#facc15"
	}
	if r > 1200 && r < 1601 {
		// return "rgba(103, 190, 51, 0.7)"
		return "#22c55e"
	}
	if r > 1600 && r < 2001 {
		// return "rgba(106, 236, 255, 0.7)"
		return "#14b8a6"
	}
	if r > 2000 && r < 2451 {
		// return "rgba(46, 174, 193, 0.7)"
		return "#3b82f6"
	}
	if r > 2450 && r < 2901 {
		// return "rgba(208, 108, 255, 0.7)"
		return "#9333ea"
	}
	if r > 2900 {
		// return "rgba(142, 65, 177, 0.7)"
		return "#a855f7"
	}
	return "#64748b"
}
