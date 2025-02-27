package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a) // this will print [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set:", a)    // this will print [0 0 0 0 100]
	fmt.Println("get:", a[4]) // this will print 100

	fmt.Println("len:", len(a)) // this will print 5
	fmt.Println("cap:", cap(a)) // this will print 5

	// declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array b:", b) // this will print [1 2 3 4 5]

	// use compiler to count the length of the array
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("array c:", c) // this will print [1 2 3 4 5]

	// specify the index with :, the elements in between will be initialized with 0
	d := [5]int{1: 10, 3: 30}
	fmt.Println("array d:", d) // this will print [0 10 0 30 0]

	// you can create 2D arrays -> like C
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D array:", twoD) // this will print [[0 1 2] [1 2 3]]

	// or initialize it in one line
	twoD2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D array 2:", twoD2) // this will print [[1 2 3] [4 5 6]]
}
