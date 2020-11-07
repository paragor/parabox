package main

import (
	"fmt"
)

func main() {
	out := Output{}
	out.Start()

	out.End()
}

type Offer struct {
	OfferId   string `json:"offer_id"`
	MarketSku uint32 `json:"market_sku"`
	Price     uint16 `json:"price"`
}

func offersTo()  {

}

type Output struct {
	isDirty bool
}

func (o Output) Add(offers string) {
	if o.isDirty && len(offers) > 0 {
		fmt.Printf(",")
	}
	fmt.Println(offers)
	o.isDirty = true
}

func (o Output) Start() {
	fmt.Printf("{\"offers\":[")
}
func (o Output) End() {
	fmt.Printf("]}")
}
