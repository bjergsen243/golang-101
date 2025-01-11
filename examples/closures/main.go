package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"John", "Doe", 25},
		{"Jane", "Smith", 20},
		{"Mike", "Johnson", 30},
	}

	fmt.Println("People:", people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println("People after sort:", people)

	// closure example
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
