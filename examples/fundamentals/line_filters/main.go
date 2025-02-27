package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cat file_name | go run examples/line_filters/main.go

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
