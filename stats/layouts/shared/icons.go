package shared

import (
	"fmt"
)

var (
	IconDirectionUpSmall   = "upSmall"
	IconDirectionUpLarge   = "upLarge"
	IconDirectionDownSmall = "downSmall"
	IconDirectionDownLarge = "downLarge"
	IconSizeVariations     = []string{IconDirectionUpSmall, IconDirectionUpLarge, IconDirectionDownSmall, IconDirectionDownLarge}

	IconDirectionHorizontal = "horizontal"
	IconDirectionVertical   = "vertical"
)

var IconCircle = "circle"
var IconsLines = make(map[string]string)
var IconsArrows = make(map[string]string)
var IconsRatingWithLevels = make(map[int]map[string]string)

func init() {
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
