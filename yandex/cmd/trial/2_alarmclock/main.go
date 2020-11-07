package main

import (
	"fmt"
	"sort"
)

func main() {
	//var N int
	//var X int
	//var K int
	//
	//fmt.Scan(&N, &X, &K)
	//var clocks []int
	//var next int
	//for i := 0; i < N; i++ {
	//	fmt.Scan(&next)
	//	clocks = append(clocks, next)
	//}
	//
	//fmt.Println(FindAlarms(clocks, X, K))

	fmt.Println(FindAlarms([]int{5, 22, 17, 13, 8, 100,100}, 7, 12))
	//fmt.Println(FindAlarms([]int{1, 2, 3, 4, 5,6}, 5, 10))
	//fmt.Println(FindAlarms([]int{1, 100}, 5, 3))
}

type Window struct {
	w []int
}

func (window *Window) Next(delta int) int {
	next := window.w[0] + delta
	window.w[0] = next
	sort.Ints(window.w)

	return next
}

func (window *Window) TryNext(delta int) int {
	return window.w[0] + delta
}

func FindAlarms(clocks []int, delta int, clockRequired int) (time int) {
	sort.Ints(clocks)

	window := Window{w: append([]int{}, clocks...)}
	var result []int
	for _, c := range clocks {
		for window.TryNext(delta) < c && len(result) < clockRequired {
			next := window.Next(delta)
			if len(result) > 0 && next == result[len(result)-1] {
				continue
			}
			result = append(result, next)
		}
		result = append(result, c)
	}
	for len(result) < clockRequired {
		next := window.Next(delta)
		if len(result) > 0 && next == result[len(result)-1] {
			continue
		}
		result = append(result, next)
	}

	return result[clockRequired-1]
}
