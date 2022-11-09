package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	x := os.Args
	lengthOfArg := len(x) - 1
	if isEven(lengthOfArg) == true {
		EvenMsg := "I have an even number of arguments"
		printStr(EvenMsg)
	} else {
		OddMsg := "I have an odd number of arguments"
		printStr(OddMsg)
	}
}
