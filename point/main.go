package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbr(n int) {
	const Zero = 48
	sign := 1
	if n == 0 {
		z01.PrintRune(rune(Zero))
	}
	if n < 0 {
		z01.PrintRune(45)
		sign = -1
	}
	var k []int
	for n != 0 {

		a := n % 10 * sign
		n = n / 10
		k = append(k, a)

	}
	for i := len(k) - 1; i >= 0; i-- {
		z01.PrintRune(rune(Zero + k[i]))
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
