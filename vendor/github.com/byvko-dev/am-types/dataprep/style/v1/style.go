package style

import (
	"image/color"
	"reflect"
)

// Block content alignment
type AlignItemsValue string
type JustifyContentValue string

const (
	AlignItemsVertical   AlignItemsValue = "vertical"
	AlignItemsHorizontal AlignItemsValue = "horizontal"

	JustifyContentStart        JustifyContentValue = "start"
	JustifyContentCenter       JustifyContentValue = "center"
	JustifyContentEnd          JustifyContentValue = "end"
	JustifyContentSpaceBetween JustifyContentValue = "space-between"
)

type Style struct {
	Color               color.RGBA `json:"color" bson:"color"`
	BackgroundColor     color.RGBA `json:"backgroundColor" bson:"backgroundColor"`
	BackgroundImage     string     `json:"backgroundImage" bson:"backgroundImage"`
	BackgroundImageBlur float64    `json:"backgroundImageBlur" bson:"backgroundImageBlur"`
	Invisible           bool       `json:"invisible" bson:"invisible"`

	// Font
	FontName string  `json:"fontFamily" bson:"fontFamily"`
	FontSize float64 `json:"fontSize" bson:"fontSize"`

	// Border
	BorderRadius float64 `json:"borderRadius" bson:"borderRadius"`

	// Alignment
	JustifyContent JustifyContentValue `json:"justifyContent" bson:"justifyContent"`
	AlignItems     AlignItemsValue     `json:"alignItems" bson:"alignItems"`
	Gap            float64             `json:"gap" bson:"gap"`

	// Spacing / Sizing
	PaddingLeft   float64 `json:"paddingLeft" bson:"paddingLeft"`
	PaddingRight  float64 `json:"paddingRight" bson:"paddingRight"`
	PaddingTop    float64 `json:"paddingTop" bson:"paddingTop"`
	PaddingBottom float64 `json:"paddingBottom" bson:"paddingBottom"`
	GrowX         bool    `json:"growX" bson:"growX"`
	GrowY         bool    `json:"growY" bson:"growY"`

	Width    float64 `json:"width" bson:"width"`
	MaxWidth float64 `json:"maxWidth" bson:"maxWidth"`
	MinWidth float64 `json:"minWidth" bson:"minWidth"`
}

// Merge o into s, overwriting s values with o values if they are not nil
func (s Style) Merge(o Style) Style {
	if reflect.TypeOf(o) != reflect.TypeOf(s) {
		return s
	}
	var result Style
	for i := 0; i < reflect.TypeOf(o).NumField(); i++ {
		// handler slices
		if reflect.TypeOf(o).Field(i).Type.Kind() == reflect.Slice {
			if !reflect.ValueOf(o).Field(i).IsNil() {
				reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(o).Field(i))
			} else {
				reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(s).Field(i))
			}
			continue
		}

		if reflect.ValueOf(o).Field(i).Interface() != reflect.Zero(reflect.TypeOf(o).Field(i).Type).Interface() {
			reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(o).Field(i))
		} else {
			reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(s).Field(i))
		}
	}
	return result
}

const (
	IconPositionLeft = iota
	IconPositionRight
)
