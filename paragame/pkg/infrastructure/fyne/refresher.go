package fyne

import (
	"fyne.io/fyne/canvas"
	"github.com/nfnt/resize"
	"github.com/paragor/parabox/paragame/pkg/domain/core"
	"image"
)

type ImageRefresher struct {
	raster   *canvas.Raster
	imgProxy *ImageProxy
	img      image.Image
	scale    int
}

func CreateRefresher(img image.Image, scale int) (*ImageRefresher, *canvas.Raster) {
	proxy := NewImageProxy(img)
	raster := canvas.NewRasterFromImage(proxy)
	return &ImageRefresher{raster: raster, img: img, imgProxy: proxy, scale: scale}, raster
}

func (a *ImageRefresher) NotifyFieldChanged(field core.Field) {
	a.imgProxy.Update(
		resize.Resize(
			uint(a.raster.Size().Width*a.scale),
			uint(a.raster.Size().Height*a.scale),
			a.img,
			resize.NearestNeighbor,
		),
	)
	a.raster.Refresh()
}
