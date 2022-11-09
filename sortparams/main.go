package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	compA := ""
	compB := ""
	for i := 1; i < len(x); i++ {
		for i := 1; i < len(x)-1; i++ {
			word1 := []rune(x[i])
			word2 := []rune(x[i+1])
			if word1[0] > word2[0] {
				compA = x[i]
				compB = x[i+1]
				x[i] = compB
				x[i+1] = compA
			}
		}
	}
	for i := 1; i < len(x); i++ {
		for _, y := range x[i] {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
