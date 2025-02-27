package main

// import "embed"
//
//go:embed examples/embed_directive/main.go
var fileString string

func main() {
	print(fileString)
}
