package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	d := len(x) - 1
	for argnum := d; argnum > 0; argnum-- {
		for argelem := 0; argelem < 1; argelem++ {
			for _, y := range x[argnum] {
				argelem = argelem + 1
				z01.PrintRune(y)
			}
			z01.PrintRune('\n')
		}
	}
}
