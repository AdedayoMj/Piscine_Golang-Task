package main

import (
	"github.com/01-edu/z01"
)

// Create a .go file.
// The code below has to be copied in that file.
// The necessary changes have to be applied so that the program works.
// The program must be submitted inside a folder with the name point.

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbr(n int) {
	sign := 1
	if n < 0 {
		sign = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * sign
		if f != 0 {
			PrintNbr(f)
		}
		k := (n % 10 * sign) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(", ")
	PrintStr("y = ")
	PrintNbr(points.y)
	z01.PrintRune(rune('\n'))
}
