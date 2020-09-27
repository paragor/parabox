package main

import (
	"github.com/paragor/parabox/paragame/pkg/domain/core"
	"github.com/paragor/parabox/paragame/pkg/domain/game_of_life"
	"github.com/paragor/parabox/paragame/pkg/domain/random_life"
	"github.com/paragor/parabox/paragame/pkg/infrastructure/cli"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	field := core.NewStdField(100, 20)
	random_life.SeedField(field, 50, game_of_life.LifeCell)

	output := cli.NewFieldOutput(cli.SymbolMap{
		game_of_life.EmptyCell: " ",
		game_of_life.LifeCell:  "+",
	})
	output.NotifyFieldChanged(field)

	game := core.NewGame(
		field,
		core.NewTimeTicker(time.Millisecond * 100),
		game_of_life.NewGameOfLifeMutator(),
		output,
	)
	game.Start()
	<-make(chan bool)
}

