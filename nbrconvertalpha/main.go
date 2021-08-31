package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	mString := []rune(s)
	const Zero = 48
	res := 0
	length := len([]rune(s))

	if length == 0 {
		return 0
	}
	var sign bool
	if mString[0] == 45 {
		mString[0] = 48
		sign = true
	} else {
		if mString[0] == 43 {
			mString[0] = 48
		}
	}
	for _, v := range mString {
		if v < 48 || v > 57 {
			return 0
		}
	}

	for _, v := range mString {
		a := int(v - Zero)
		res = res*10 + a
	}
	if sign {
		res = res * -1
	}

	return res
}

func main() {
	nbr := os.Args[1:]
	s := "abcdefghijklmnopqrstuvwxyz"
	m := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alph := []rune(s)
	upAlph := []rune(m)
	flag := false

	if len(nbr) < 1 {
		return
	}
	if nbr[0] == "--upper" {
		flag = true
		nbr = nbr[1:]
	}
	for _, n := range nbr {
		in := Atoi(n)
		if in < 1 || in > 26 {
			z01.PrintRune(' ')
			continue
		}
		if flag {
			z01.PrintRune(upAlph[in-1])
		}
		if !flag {
			z01.PrintRune(alph[in-1])
		}
	}
	z01.PrintRune('\n')
}
