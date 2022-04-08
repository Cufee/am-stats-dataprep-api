package icons

import (
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

	IconDirectionUp    = "up"
	IconDirectionDown  = "down"
	IconDirectionLeft  = "left"
	IconDirectionRight = "right"
)

var IconsLines = make(map[string]string)
var IconsCircle = make(map[string]string)
var IconsArrows = make(map[string]string)
var IconsTriangles = make(map[string]string)

func init() {
	IconsCircle[""] = "circle"
	IconsCircle[IconDirectionUp] = "circle"
	IconsCircle[IconDirectionDown] = "circle"
	IconsCircle[IconDirectionLeft] = "circle"
	IconsCircle[IconDirectionRight] = "circle"

	IconsArrows[""] = ""
	IconsArrows[IconDirectionUp] = "arrowUp"
	IconsArrows[IconDirectionDown] = "arrowDown"
	IconsArrows[IconDirectionLeft] = "arrowLeft"
	IconsArrows[IconDirectionRight] = "arrowRight"

	IconsLines[""] = ""
	IconsLines[IconDirectionDown] = "lineVertical"
	IconsLines[IconDirectionRight] = "lineVertical"
	IconsLines[IconDirectionLeft] = "lineHorizontal"
	IconsLines[IconDirectionUp] = "lineHorizontal"

	IconsTriangles[""] = ""
	IconsTriangles[IconDirectionUp] = "triangleUp"
	IconsTriangles[IconDirectionDown] = "triangleDown"
	IconsTriangles[IconDirectionLeft] = "triangleLeft"
	IconsTriangles[IconDirectionRight] = "triangleRight"
}
