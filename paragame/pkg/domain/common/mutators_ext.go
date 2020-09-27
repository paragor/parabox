package common

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
)

type mutatorFuncAdapter struct {
	fn func(field core.Field, x int, y int)
}

func (m *mutatorFuncAdapter) Mutate(field core.Field, x int, y int) {
	m.fn(field, x, y)
}

func MutatorFunc(mutator func(field core.Field, x int, y int)) core.Mutator {
	return &mutatorFuncAdapter{fn: mutator}
}

// -----------

type MultiMutator struct {
	mutators []core.Mutator
}

func NewMultiMutator(one core.Mutator, other ...core.Mutator) *MultiMutator {
	mutators := make([]core.Mutator, 0, 1+len(other))
	mutators = append(mutators, one)
	mutators = append(mutators, other...)
	return &MultiMutator{mutators: mutators}
}

func (mm *MultiMutator) Mutate(field core.Field, x int, y int) {
	for _, mutator := range mm.mutators {
		mutator.Mutate(field, x, y)
	}
}

// ---------

type SomeTimesMutator struct {
	counter      int
	stepInterval int
	mutator      core.Mutator
}

func (s *SomeTimesMutator) NotifyFieldChanged(field core.Field) {
	s.counter++
	if s.counter > s.stepInterval {
		s.counter = 0
	}
}

func NewSomeTimesMutator(mutator core.Mutator, stepInterval int) *SomeTimesMutator {
	return &SomeTimesMutator{counter: 0, stepInterval: stepInterval, mutator: mutator}
}

func (s *SomeTimesMutator) Mutate(field core.Field, x int, y int) {
	if s.counter == s.stepInterval {
		s.mutator.Mutate(field, x, y)
	}
}
