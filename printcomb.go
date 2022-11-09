package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var a rune = '0'
	var b rune = '0'
	var c rune = '0'
	bla := 1
	for i := 1; i < 999; i++ {
		if a == '9' {
			if b == '9' {
				c = c + 1
				b = '0'
			} else {
				b = b + 1
			}
			a = '0'
		} else {
			a = a + 1
		}

		if c < b {
			if b < a {
				if bla == 0 {

					z01.PrintRune(',')
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(b)
					z01.PrintRune(a)
				} else {
					z01.PrintRune(c)
					z01.PrintRune(b)
					z01.PrintRune(a)
					bla = 0
				}
			}
		}

	}
	z01.PrintRune('\n')
}
