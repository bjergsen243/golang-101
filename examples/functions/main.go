package main

import "fmt"

func main() {
	fmt.Println("annotated function")
	anonymousFunc()

	fmt.Println("plus function")
	resPlus := plus(1, 2)
	fmt.Println("1 + 2 =", resPlus)

	fmt.Println("return multiple values")
	a, b := returnMultipleValues()
	fmt.Println("a:", a, "b:", b)

	fmt.Println("variadic function")
	variadicFunc(1, 2)
	variadicFunc(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	variadicFunc(nums...)
}

// annonymous function is useful when you want to define a function inline
// without having to name it
func anonymousFunc() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside the anonymous function")
		}(i)
	}
}

func plus(a, b int) int {
	return a + b
}

func returnMultipleValues() (int, int) {
	return 3, 7
}

func variadicFunc(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total is:", total)
}
