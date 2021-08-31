package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	store := []rune(args[0])
	length := len(store)

	for i := 2; i < length; i++ {
		z01.PrintRune(store[i])
	}
	z01.PrintRune('\n')
}
