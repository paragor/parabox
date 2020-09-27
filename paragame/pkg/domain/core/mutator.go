package core

type Mutator interface {
	Mutate(field Field, x int, y int)
}
