package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// print Go values
	fmt.Printf("struct1: %v\n", p)

	// include struct field name
	fmt.Printf("struct2: %+v\n", p)

	// print Go syntax representation of the value
	fmt.Printf("struct3: %#v\n", p)

	// print type of value
	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 14.56)

	fmt.Printf("float2: %e\n", 14344.0)

	fmt.Printf("float3: %E\n", 14322.0)

	fmt.Printf("string1: %s\n", "\"string\"")

	fmt.Printf("string2: %q\n", "\"string\"")

	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 34.5)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width3: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
