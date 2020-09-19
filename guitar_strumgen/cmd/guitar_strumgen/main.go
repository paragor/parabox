package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func main() {

	strum := NewStrum()
	for i := 0; i < 6; i++ {
		for strum.Add(RandomSymbol()) != nil {
		}
	}

	fmt.Println(strum.String())
}

func RandomSymbol() symbol {
	symbols := []symbol{Up, Down, Empty}
	return symbols[rand.Intn(len(symbols))]
}

type symbol string

const (
	Up    symbol = "↑"
	Down  symbol = "↓"
	Empty symbol = "·"
)
var (
	NotCorrectSymbolError = errors.New("not correct new symbol")
)

type Strum struct {
	symbols []symbol
}

func NewStrum() *Strum {
	return &Strum{symbols: make([]symbol, 0)}
}

func (strum *Strum) Add(newSymbol symbol) error {
	if len(strum.symbols) == 0 {
		strum.symbols = append(strum.symbols, newSymbol)
		return nil
	}

	last := strum.symbols[len(strum.symbols)-1]

	if last == Empty {
		strum.symbols = append(strum.symbols, newSymbol)
		return nil
	}
	if last == Up && newSymbol == Up {
		return NotCorrectSymbolError
	}
	if last == Down && newSymbol == Down {
		return NotCorrectSymbolError
	}

	strum.symbols = append(strum.symbols, newSymbol)
	return nil
}

func (strum *Strum) String() string {
	result := ""
	for _, s := range strum.symbols {
		result += string(s)
	}
	return result
}
