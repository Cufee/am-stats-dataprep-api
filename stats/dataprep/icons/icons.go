package icons

import (
	"image/color"
	"strconv"
)

var (
	IconColorWhite   = hexToColor("#FFFFFF")
	IconColorBlack   = hexToColor("#000000")
	IconColorRed     = hexToColor("#fb7185")
	IconColorBlue    = hexToColor("#7dd3fc")
	IconColorTeal    = hexToColor("#2dd4bf")
	IconColorGreen   = hexToColor("#34d399")
	IconColorYellow  = hexToColor("#fcfc7d")
	IconColorPurple  = hexToColor("#c084fc")
	IconColorNeutral = hexToColor("#94a3b8")

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

func hexToColor(hex string) color.RGBA {
	var r, g, b, a int64 = 0, 0, 0, 255
	if len(hex) == 7 {
		r, _ = strconv.ParseInt(hex[1:3], 16, 0)
		g, _ = strconv.ParseInt(hex[3:5], 16, 0)
		b, _ = strconv.ParseInt(hex[5:7], 16, 0)
	} else if len(hex) == 9 {
		r, _ = strconv.ParseInt(hex[1:3], 16, 0)
		g, _ = strconv.ParseInt(hex[3:5], 16, 0)
		b, _ = strconv.ParseInt(hex[5:7], 16, 0)
		a, _ = strconv.ParseInt(hex[7:9], 16, 0)
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}
