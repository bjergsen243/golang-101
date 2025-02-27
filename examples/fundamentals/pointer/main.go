package main

import "fmt"

func main() {
	var f *int
	failedUpdate(f)
	fmt.Println(f)

	x := 10
	update(&x)
	fmt.Println(x)

	y := 10
	failedUpdate(&y)
	fmt.Println(y)
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(px *int) {
	*px = 20
}
