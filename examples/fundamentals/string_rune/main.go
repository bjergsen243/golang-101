package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// string is a read-only slice of bytes
	stringFunc()
	// rune is an integer that represents a Unicode code point

}

func stringFunc() {
	const s = "สวัสดี" // means hello in Thai
	fmt.Println("Length of s:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // prints the hex value of the byte
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// decode each rune
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString((s[i:]))
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
