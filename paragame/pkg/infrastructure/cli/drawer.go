package cli

import (
	"fmt"
	"github.com/paragor/parabox/paragame/pkg/domain/core"
)

type SymbolMap map[core.CellState]string

func (s SymbolMap) Get(state core.CellState) string {
	str, ok := s[state]
	if !ok {
		return "?"
	}
	return str
}

type FieldCliDrawer struct {
	smap SymbolMap
}

func NewFieldOutput(symbolMap SymbolMap) *FieldCliDrawer {
	return &FieldCliDrawer{smap: symbolMap}
}

func (f *FieldCliDrawer) NotifyFieldChanged(field core.Field) {
	f.clear()
	for y := 0; y < field.MaxY(); y++ {
		row := ""
		for x := 0; x < field.MaxX(); x++ {
			row += f.smap.Get(field.Get(x, y))
		}
		fmt.Println(row)
	}
	fmt.Println()
}

func (f *FieldCliDrawer) clear() {
	fmt.Println("\033[2J")
}
