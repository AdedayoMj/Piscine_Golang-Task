package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	/*The following is a triple loop*/
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				// 		/*Ensure that i, j, and k are different from each other*/
				if i != k && i != j && j != k && i < j && j < k {
					if !(i == '7' && j == '8' && k == '9') {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
