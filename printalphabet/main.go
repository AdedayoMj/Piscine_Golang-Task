package main

import "github.com/01-edu/z01"

func main() {
	for ch := 97; ch <= 122; ch++ {
		z01.PrintRune(rune(ch))
	}
	z01.PrintRune('\n')
}
