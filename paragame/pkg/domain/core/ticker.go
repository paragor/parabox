package core

import (
	"time"
)

type Ticker interface {
	Tick() <-chan struct{}
}

type TimeTicker struct {
	ticker *time.Ticker
	c      <-chan struct{}
}

func NewTimeTicker(interval time.Duration) *TimeTicker {
	ticker := time.NewTicker(interval)
	c := make(chan struct{})
	go func() {
		for {
			<-ticker.C
			c <- struct{}{}
		}
	}()
	return &TimeTicker{ticker: ticker, c: c}
}

func (t *TimeTicker) Tick() <-chan struct{} {
	return t.c
}
