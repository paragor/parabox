package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["azaza"] = 33
	val, ok := m["azazaa"]
	fmt.Println(val, ok)

}
