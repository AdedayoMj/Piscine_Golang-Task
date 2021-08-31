package main

// Create a .go file.
// The code below has to be copied in that file.
// The necessary changes have to be applied so that the program works.
// The program must be submitted inside a folder with the name boolean.

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayStr := []rune(s)
	for _, r := range arrayStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) == true {
		return true
	} else {
		return false
	}
}

func even(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	args := os.Args

	lengthOfArg := 0
	for range args {
		lengthOfArg++
	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg + 1) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
