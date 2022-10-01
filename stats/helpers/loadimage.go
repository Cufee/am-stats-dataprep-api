package helpers

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

func LoadImage(path string) (image.Image, error) {
	return gg.LoadImage(path)
}

func BlurImage(img image.Image, blur float64) image.Image {
	return imaging.Blur(img, blur)
}
