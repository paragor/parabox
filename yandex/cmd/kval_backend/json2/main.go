package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var count int
	_, err := fmt.Scanln(&count)
	if err != nil {
		panic(err)
	}

	out := Output{}
	out.Start()
	offers := struct {
		Offers []Offer `json:"offers"`
	}{Offers: nil}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()

	for i := 0; i < count && scanner.Scan(); i++ {
		err = json.Unmarshal(scanner.Bytes(), &offers)
		if err != nil {
			panic(err)
		}
		out.Add(offers.Offers)

		offers.Offers = nil
	}

	out.End()
}

type Offer struct {
	OfferId   string `json:"offer_id"`
	MarketSku int `json:"market_sku"`
	Price     int `json:"price"`
}

type Output struct {
	isDirty bool
}

func (o *Output) Add(offers []Offer) {
	if len(offers) == 0 {
		return
	}
	for _, offer := range offers {
		b, err := json.Marshal(offer)
		if err != nil {
			panic(err)
		}
		if o.isDirty {
			fmt.Print(",")
		}
		fmt.Print(string(b))
		o.isDirty = true
	}
}

func (o *Output) Start() {
	fmt.Print("{\"offers\":[")
	o.isDirty = false
}
func (o *Output) End() {
	fmt.Print("]}")
}
