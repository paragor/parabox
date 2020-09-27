package random_life

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
)

func SeedField(field core.Field, lifeCoef int, lifeCell core.CellState) {
	core.ApplyMutator(field, NewRandMutator(lifeCoef, lifeCell))
	field.Next()
}
