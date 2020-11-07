package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func main() {
	var url string
	var port int
	var a int
	var b int
	_, err := fmt.Scan(&url, &port, &a, &b)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient

	response, err := client.Get(fmt.Sprintf("%s:%d/?a=%d&b=%d", url, port, a, b))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var ints []int
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&ints)
	if err != nil {
		panic(err)
	}

	ints = FilterAndSort(ints)

	for _, val := range ints {
		fmt.Println(val)
	}
}

func FilterAndSort(ints []int) []int {
	result := make([]int, 0, len(ints))

	for _, val := range ints {
		if val < 0 {
			result = append(result, val)
		}
	}

	sort.Ints(result)

	return result
}
