package img

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
	"image"
	"image/color"
)

type SymbolMap map[core.CellState]color.RGBA

func (s SymbolMap) Get(state core.CellState) color.RGBA {
	str, ok := s[state]
	if !ok {
		return color.RGBA{}
	}
	return str
}

// ужас конечно
type colorMap struct {
	cmap [][]color.Color
	rect image.Rectangle
}

func (cm *colorMap) ColorModel() color.Model {
	return color.RGBAModel
}

func (cm *colorMap) Bounds() image.Rectangle {
	return cm.rect
}

func (cm *colorMap) At(x, y int) color.Color {
	if cm.cmap == nil {
		return color.RGBA{}
	}
	return cm.cmap[y][x]
}
func (cm *colorMap) update(field core.Field, symbolMap SymbolMap) {
	cmap := make([][]color.Color, field.MaxY())
	for y := 0; y < field.MaxY(); y++ {
		cmap[y] = make([]color.Color, field.MaxX())
		for x := 0; x < field.MaxX(); x++ {
			cmap[y][x] = symbolMap.Get(field.Get(x, y))
		}
	}
	cm.cmap = cmap
	cm.rect = image.Rectangle{
		Max: struct{ X, Y int }{
			X: field.MaxX(),
			Y: field.MaxY(),
		},
	}
}

type FieldImageDrawer struct {
	smap SymbolMap
	img  colorMap
}

func NewFieldImageDrawer(smap SymbolMap) *FieldImageDrawer {
	return &FieldImageDrawer{smap: smap}
}

func (f *FieldImageDrawer) NotifyFieldChanged(field core.Field) {
	f.img.update(field, f.smap)
}

func (f *FieldImageDrawer) GetImage() image.Image {
	return &f.img
}
