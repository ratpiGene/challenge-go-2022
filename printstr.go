package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i])) // print le caractère de la rune en pos i du string
	}
}
