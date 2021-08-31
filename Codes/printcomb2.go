package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	/*The following is a triple loop*/
	/* Using ascii*/
	// for i := 48; i <= 57; i++ {
	// 	for j := 48; j <= 57; j++ {
	// 		for k := 48; k <= 57; k++ {
	// 			for l := 48; l <= 57; l++ {
	// 				if i == k && j < l || i < k {
	// 					if !(i == 57 && j == 56 && k == 57 && l == 57) {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i == k && j < l || i < k {
						if !(i == '9' && j == '8' && k == '9' && l == '9') {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(' ')
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(' ')
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
						}
					}
				}
			}
		}
	}

	z01.PrintRune('\n')
}
