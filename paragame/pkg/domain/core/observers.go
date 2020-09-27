package core

type FieldChangeObserver interface {
	NotifyFieldChanged(field Field)
}

type NullFieldChangedObserver struct{}

func (NullFieldChangedObserver) NotifyFieldChanged(field Field) {}
