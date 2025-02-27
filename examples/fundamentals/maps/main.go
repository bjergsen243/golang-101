package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	v2 := m["k2"]
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("after delete k2:", m)

	clear(m)
	fmt.Println("after clear:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n2)

	if maps.Equal(n, n2) {
		fmt.Println("n and n2 are equal")
	}

}
