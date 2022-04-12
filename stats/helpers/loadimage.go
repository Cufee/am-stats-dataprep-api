package helpers

import (
	"image"

	"github.com/fogleman/gg"
)

func LoadImage(path string) (image.Image, error) {
	return gg.LoadImage(path)
}
