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

