package common

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
)

func CalculateAroundCells(field core.Field, x, y int) map[core.CellState]int {
	cellsMap := make(map[core.CellState]int)
	for dy := y - 1; dy <= y+1; dy++ {
		if dy >= field.MaxY() || dy < 0 {
			continue
		}
		for dx := x - 1; dx <= x+1; dx++ {
			if dx >= field.MaxX() || dx < 0 {
				continue
			}
			if dx == x && dy == y {
				continue
			}
			cell := field.Get(dx, dy)
			count, ok := cellsMap[cell]
			if !ok {
				count = 0
			}
			cellsMap[cell] = count + 1
		}

	}
	return cellsMap
}
