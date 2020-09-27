package game_of_life

import (
	"github.com/paragor/parabox/paragame/pkg/domain/common"
	"github.com/paragor/parabox/paragame/pkg/domain/core"
)

const (
	EmptyCell = core.EmptyCell
	LifeCell  = core.EmptyCell + 1
)

type GameOfLifeMutator struct {
}

func NewGameOfLifeMutator() *GameOfLifeMutator {
	return &GameOfLifeMutator{}
}

func (g *GameOfLifeMutator) Mutate(field core.Field, x int, y int) {
	cellsMap := common.CalculateAroundCells(field, x, y)
	lifeCellsCount, ok := cellsMap[LifeCell]
	if !ok {
		field.Set(x, y, EmptyCell)
		return
	}

	curIsLife := field.Get(x, y) == LifeCell

	if !curIsLife && lifeCellsCount == 3 {
		field.Set(x, y, LifeCell)
		return
	}
	if curIsLife && (lifeCellsCount == 2 || lifeCellsCount == 3) {
		field.Set(x, y, LifeCell)
		return
	}

	field.Set(x, y, EmptyCell)
}

