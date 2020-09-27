package random_life

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
	"math/rand"
)

type RandMutator struct {
	lifeCoef int
	lifeCell core.CellState
}

// lifeCoef > 0 && lifeCoef <= 100
func NewRandMutator(lifeCoef int, lifeCell core.CellState) *RandMutator {
	if lifeCoef < 0 || lifeCell > 100 {
		panic("life coef is invalid")
	}
	return &RandMutator{lifeCoef: lifeCoef, lifeCell: lifeCell}
}

func (r RandMutator) Mutate(field core.Field, x int, y int) {
	if rand.Intn(100) > r.lifeCoef && field.Get(x, y) == core.EmptyCell {
		field.Set(x, y, r.lifeCell)
	}
}
