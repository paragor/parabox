package fyne

import (
	"image"
	"image/color"
)

// Позволяет подменять настоящее изображение
// мне было лень писать свой Raster, поэтому заюзал имеющийся + этот костылек
type ImageProxy struct {
	img image.Image
}


func (ip *ImageProxy) ColorModel() color.Model {
	return ip.img.ColorModel()
}

func (ip *ImageProxy) Bounds() image.Rectangle {
	return ip.img.Bounds()
}

func (ip *ImageProxy) At(x, y int) color.Color {
	return ip.img.At(x, y)
}

func NewImageProxy(img image.Image) *ImageProxy {
	return &ImageProxy{img: img}
}

func (ip *ImageProxy) Update(img image.Image) {
	ip.img = img
}
