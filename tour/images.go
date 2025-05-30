package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 10, 10)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	v := uint8((x * y) % 255)
	return color.RGBA{v, v, v, 255}
}

func main() {
	m := &Image{}
	pic.ShowImage(m)
}
