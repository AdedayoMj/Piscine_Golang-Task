package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var num [9]rune
	var i int
	var counter int
	if n > 9 {
		n = 9
	}
	if n == 0 {
		z01.PrintRune('\n')
		return
	}
	for i = 0; i < n; i++ {
		num[i] = rune(i + 48)
	}
	for {
		for i = 0; i < n; i++ {
			z01.PrintRune(num[i])
		}
		i = n - 1
		for counter = 0; true; i-- {
			num[i]++
			if num[i] < rune('9'-n+i+2) {
				break
			}
			if i == 0 {
				z01.PrintRune('\n')
				return
			}
			counter++
		}
		for ; counter != 0; counter-- {
			num[n-counter] = num[n-counter-1] + 1
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
