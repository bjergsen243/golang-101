package main

import (
	"fmt"
	"os"
)

func main() {
	// defer won't run when using exit
	defer fmt.Println("!")

	os.Exit(3)
}
