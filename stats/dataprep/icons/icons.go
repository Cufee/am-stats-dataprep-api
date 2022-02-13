package icons

const (
	IconColorWhite   = "white"
	IconColorBlack   = "black"
	IconColorRed     = "#fb7185"
	IconColorBlue    = "#7dd3fc"
	IconColorTeal    = "#2dd4bf"
	IconColorGreen   = "#34d399"
	IconColorYellow  = "#fde047"
	IconColorPurple  = "#c084fc"
	IconColorNeutral = "#94a3b8"

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
