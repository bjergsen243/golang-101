package main

import (
	"flag"
	"fmt"
)

// ./main -word=opt -numb=7 -fork -svar=flag
// ./main -word=opt
// ./main -word=opt a1 a2 a3
// ./main -word=opt a1 a2 a3 -numb=7
// ./main -h
// ./main -wat

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("numb:", *wordPtr)
	fmt.Println("word:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
