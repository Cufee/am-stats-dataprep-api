package icons

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/helpers"
)

var (
	IconColorWhite   = helpers.HexToColor("#FFFFFF")
	IconColorBlack   = helpers.HexToColor("#000000")
	IconColorRed     = helpers.HexToColor("#fb7185")
	IconColorBlue    = helpers.HexToColor("#7dd3fc")
	IconColorTeal    = helpers.HexToColor("#2bcee3")
	IconColorGreen   = helpers.HexToColor("#34d399")
	IconColorYellow  = helpers.HexToColor("#fcfc7d")
	IconColorPurple  = helpers.HexToColor("#c084fc")
	IconColorNeutral = helpers.HexToColor("#94a3b8")

	IconDirectionUpSmall   = "upSmall"
	IconDirectionUpLarge   = "upLarge"
	IconDirectionDownSmall = "downSmall"
	IconDirectionDownLarge = "downLarge"
	IconSizeVariations     = []string{IconDirectionUpSmall, IconDirectionUpLarge, IconDirectionDownSmall, IconDirectionDownLarge}

	IconDirectionHorizontal = "horizontal"
	IconDirectionVertical   = "vertical"
)

var IconsLines = make(map[string]string)
var IconsCircle = make(map[string]string)
var IconsArrows = make(map[string]string)
var IconsRatingWithLevels = make(map[int]map[string]string)

func init() {
	IconsCircle[IconDirectionUpSmall] = "circle"
	IconsCircle[IconDirectionUpLarge] = "circle"
	IconsCircle[IconDirectionDownLarge] = "circle"
	IconsCircle[IconDirectionDownSmall] = "circle"

	IconsArrows[IconDirectionUpSmall] = "singleArrowUp"
	IconsArrows[IconDirectionUpLarge] = "doubleArrowUp"
	IconsArrows[IconDirectionDownLarge] = "doubleArrowDown"
	IconsArrows[IconDirectionDownSmall] = "singleArrowDown"

	IconsLines[IconDirectionVertical] = "lineVertical"
	IconsLines[IconDirectionHorizontal] = "lineHorizontal"

	for level := 1; level <= 3; level++ {
		IconsRatingWithLevels[level] = make(map[string]string)
		for _, size := range IconSizeVariations {
			IconsRatingWithLevels[level][size] = "rating-" + fmt.Sprint(level)
		}
	}
}

func IconsRating(level int) map[string]string {
	return IconsRatingWithLevels[level]
}
