package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("length of s", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		fmt.Printf("%T", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count", utf8.RuneCountInString(s))

	for index, runevalue := range s {
		fmt.Printf("%#U starts at %d\n", runevalue, index)
	}
}
