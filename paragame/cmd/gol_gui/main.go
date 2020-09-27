package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/paragor/parabox/paragame/pkg/domain/common"
	"github.com/paragor/parabox/paragame/pkg/domain/core"
	"github.com/paragor/parabox/paragame/pkg/domain/game_of_life"
	"github.com/paragor/parabox/paragame/pkg/domain/random_life"
	infr_fyne "github.com/paragor/parabox/paragame/pkg/infrastructure/fyne"
	"github.com/paragor/parabox/paragame/pkg/infrastructure/img"
	"image/color"
	"math"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	WINDOW_X = 260
	WINDOW_Y = 160

	SCALE = 2
)

var FIELD_X = 256
var FIELD_Y = int(math.Floor(float64(WINDOW_Y) / float64(WINDOW_X) * float64(FIELD_X)))

func main() {
	field := core.NewStdField(FIELD_X, FIELD_Y)
	random_life.SeedField(field, 50, game_of_life.LifeCell)

	mutator := common.NewMultiMutator(
		game_of_life.NewGameOfLifeMutator(),
	)

	imgOutput := img.NewFieldImageDrawer(img.SymbolMap{
		game_of_life.EmptyCell: color.RGBA{},
		game_of_life.LifeCell:  color.RGBA{R: 255, G: 255, B: 255, A: 125},
	})
	imgOutput.NotifyFieldChanged(field)

	a := app.New()
	window := a.NewWindow("Game of Life [ext]")
	fyneRefresher, raster := infr_fyne.CreateRefresher(imgOutput.GetImage(), SCALE)
	window.SetContent(raster)

	observer := common.NewMultiObserver(
		imgOutput,
		fyneRefresher,
	)

	game := core.NewGame(
		field,
		core.NewTimeTicker(time.Second/60),
		mutator,
		observer,
	)

	game.Start()
	defer game.Stop()
	//todo сделать локи на мутирование и старт-стоп
	isRunning := true
	window.Resize(fyne.NewSize(WINDOW_X*SCALE+10, WINDOW_Y*SCALE+10))
	window.SetOnClosed(func() {
		os.Exit(0)
	})
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeySpace {
			if isRunning {
				game.Stop()
			} else {
				game.Start()
			}
			isRunning = !isRunning
			return
		}

		if event.Name == fyne.KeyReturn {
			random_life.SeedField(field, 1, game_of_life.LifeCell)
			observer.NotifyFieldChanged(field)

			return
		}

	})
	window.ShowAndRun()
	<-make(chan bool)
}
