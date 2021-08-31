package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argIn := os.Args
	for i, s := range argIn {
		if i != 0 {
			for _, a := range s {
				z01.PrintRune(a)
			}
			z01.PrintRune('\n')
		}
	}
}
