package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	for _, y := range x[0][2:] {
		z01.PrintRune(y)
	}
	z01.PrintRune('\n')
}
