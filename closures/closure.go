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

}
