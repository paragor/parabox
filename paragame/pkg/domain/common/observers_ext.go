package common

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
)

type MultiObserver struct {
	observers []core.FieldChangeObserver
}

func (mo *MultiObserver) NotifyFieldChanged(field core.Field) {
	for _, observer := range mo.observers {
		observer.NotifyFieldChanged(field)
	}
}

func NewMultiObserver(one core.FieldChangeObserver, other ...core.FieldChangeObserver) *MultiObserver {
	observers := make([]core.FieldChangeObserver, 0, 1+len(other))
	observers = append(observers, one)
	observers = append(observers, other...)
	return &MultiObserver{observers: observers}
}

// -----

type observerFuncAdapter struct {
	fn func(field core.Field)
}

func (o *observerFuncAdapter) NotifyFieldChanged(field core.Field) {
	o.fn(field)
}

func ObserverFunc(observer func(field core.Field)) core.FieldChangeObserver {
	return &observerFuncAdapter{fn: observer}
}

// -----

type SomeTimesObserver struct {
	counter      int
	stepInterval int
	observer     core.FieldChangeObserver
}

func NewSomeTimesObserver(stepInterval int, observer core.FieldChangeObserver) *SomeTimesObserver {
	return &SomeTimesObserver{counter: 0, stepInterval: stepInterval, observer: observer}
}

func (s *SomeTimesObserver) NotifyFieldChanged(field core.Field) {
	s.counter++
	if s.counter == s.stepInterval {
		s.observer.NotifyFieldChanged(field)
		s.counter = 0
	}
}
