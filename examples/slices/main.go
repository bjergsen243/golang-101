package main

import (
	"fmt"
	"slices"
)

func main() {
	// declare a uninitialized slice
	var s []string
	fmt.Println("uninit:", s)
	fmt.Println("check nil:", s == nil)
	fmt.Println("Check len equals to 0:", len(s) == 0)

	fmt.Println("create slice with make")
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	fmt.Println("check nil:", s == nil)
	// now this slice is not nil and has a length of 3

	fmt.Println("set and get")
	// set and get
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len after set:", len(s))

	fmt.Println("append")
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s after append:", s)

	// create a slice from copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slice 2:5", l)

	l = s[:5]
	fmt.Println("slice :5", l)

	l = s[2:]
	fmt.Println("slice 2:", l)

	l[0] = "x"
	fmt.Println("s after change l[0]:", s)
	// if you change element from copied slice
	// the original slice will be changed also

	// declare and initialize a slice in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	fmt.Println("dcl 2:", t2)
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	// create a 2D slice in Go
	// is similiar to create a 2D array in C

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD slice is:", twoD)

}
