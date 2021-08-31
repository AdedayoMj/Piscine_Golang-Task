package piscine

import (
	"github.com/01-edu/z01"
)

var (
	list [100]int
	kk   int
)

func EightQueens() {
	for i := 0; i < 100; i++ {
		list[i] = 0
	}
	kk = 0
	var chessBoard [8][8]int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			chessBoard[i][j] = 0
		}
	}
	Solve8Q(chessBoard, 0)
	for i := 0; i < 100; i++ {
		for j := i; j < 100; j++ {
			if list[i] > list[j] {
				Swap(&list[i], &list[j])
			}
		}
	}
	for ii := 0; ii < 100; ii++ {
		if list[ii] != 0 {
			tmp := 10000000
			runeNbr := list[ii]
			for j := 0; j < 8; j++ {

				sym := runeNbr / tmp
				runeNbr = runeNbr - sym*tmp
				tmp = tmp / 10
				z01.PrintRune(rune(48 + sym))

			}
			z01.PrintRune('\n')
		}
	}
}

func IsSafe(tableCheck [8][8]int, x int, y int) bool {
	for i := 0; i < y; i++ {
		if tableCheck[x][i] == 1 {
			return false
		}
	}
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7; j++ {
			if tableCheck[i][j] == 1 && (i-x == j-y || i-x == y-j) {
				return false
			}
		}
	}
	return true
}

func Solve8Q(table [8][8]int, col int) bool {
	if col == 8 {
		PrintSeq(table)
		return true
	}
	res := false
	for i := 0; i < 8; i++ {
		if IsSafe(table, i, col) {
			table[i][col] = 1
			res = Solve8Q(table, col+1)
			table[i][col] = 0
		}
	}
	return res
}

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

func PrintSeq(t [8][8]int) {
	tmp := 10000000
	k := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if t[i][j] == 1 {
				k = k + (j+1)*tmp
				tmp = tmp / 10
			}
		}
	}
	kk++
	list[kk] = k
}
