package main

import "github.com/01-edu/z01"

func main() {
	for ch := 122; ch >= 97; ch-- {
		z01.PrintRune(rune(ch))
	}
	z01.PrintRune('\n')
}
