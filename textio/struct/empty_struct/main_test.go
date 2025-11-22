package emptystruct

import "fmt"

func main() {
	// take up 0 bytes of memory
	empty := struct{}{}

	fmt.Printf("Empty struct: %v", empty)

	type emptyStruct struct{}

	empty = emptyStruct{}

	fmt.Printf("Empty struct: %v", empty)
}
