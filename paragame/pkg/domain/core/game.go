package core

import (
	"errors"
	"time"
)

type Game struct {
	field                Field
	ticker               Ticker
	mutator              Mutator
	stop                 chan struct{}
	filedChangedObserver FieldChangeObserver
}

func NewGame(field Field, ticker Ticker, mutator Mutator, filedChangedObserver FieldChangeObserver) *Game {
	return &Game{
		field:                field,
		ticker:               ticker,
		mutator:              mutator,
		stop:                 make(chan struct{}),
		filedChangedObserver: filedChangedObserver,
	}
}

func (game *Game) Start() {
	go func() {
		for {
			select {
			case <-game.ticker.Tick():
				ApplyMutator(game.field, game.mutator)
				game.field.Next()
				game.filedChangedObserver.NotifyFieldChanged(game.field)
			case <-game.stop:
				return
			}
		}
	}()
}

func (game *Game) Stop() error {
	select {
	case game.stop <- struct{}{}:
		return nil
	case <-time.After(time.Second * 10):
		return errors.New("game not started")
	}

}

func ApplyMutator(field Field, mutator Mutator) {
	for x := 0; x < field.MaxX(); x++ {
		for y := 0; y < field.MaxY(); y++ {
			mutator.Mutate(field, x, y)
		}
	}
}
