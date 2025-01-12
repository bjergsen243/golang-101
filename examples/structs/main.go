package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println("create new struct:", person{"Bob", 20})

	fmt.Println("create new struct with field names:", person{name: "Alice", age: 19})

	fmt.Println("create new struct with ommited field will be zero valued:", person{name: "Fred"})

	fmt.Println("yield a pointer to struct:", &person{name: "Ann"})

	fmt.Println("idiomatic way to create new struct in constructor function:", newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println("access struct field:", s.name)

	sp := &s
	fmt.Println("access struct field via pointer:", sp.age)

	fmt.Println("structs are mutable:", s)
	sp.age = 51
	fmt.Println("structs after change value:", s)

	// we can create annonymous struct like this
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println("dog is:", dog)

}
