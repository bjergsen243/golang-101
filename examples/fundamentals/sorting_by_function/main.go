package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"beach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 23},
		{name: "Alex", age: 24},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
