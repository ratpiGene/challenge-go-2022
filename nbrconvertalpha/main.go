package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	a := 64 // resultat
	for i := 1; i < len(x); i++ {
		e := []rune(x[i])
		c := 0 // addcompteur
		d := 1 // puissance
		for i := len(e) - 1; i >= 0; i-- {
			c += int(e[i]-48) * d
			d = d * 10
		}
		if x[1] == "--upper" {
			a = 64
		} else {
			a = 96
		}
		if c >= 1 && c <= 26 {
			a = a + c
			z01.PrintRune(rune(a))
		} else {
			if i > 1 && x[1] == "--upper" {
				z01.PrintRune(' ')
			}
			if i > 0 && x[1] != "--upper" {
				z01.PrintRune(' ')
			}
		}
	}
	if len(x) > 1 {
		z01.PrintRune('\n')
	}
}
