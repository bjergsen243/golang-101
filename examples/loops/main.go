package main

import "fmt"

func main() {
	fmt.Println("basic for loop")
	basicFor()
	fmt.Println("C-like for loop")
	cLikeLoop()
	fmt.Println("Range loop")
	rangeLoop()
	fmt.Println("Infinite loop")
	infiniteLoop()
}

func basicFor() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}

func cLikeLoop() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}

func rangeLoop() {
	for i := range 3 {
		fmt.Println("range: ", i)
	}
}

func infiniteLoop() {
	for {
		fmt.Println("loop")
		break
	}
}
