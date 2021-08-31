package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	if checkInput(arguments) == true {
		table := [9][9]rune{}
		table = fillTable(table, arguments)

		if isSolve(&table) == true {
			printBorad(table)
		} else {
			fmt.Println("Error") // checkInput is not True
		}
	}
}

//print the table
func printBorad(table [9][9]rune) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if x != 8 {
				z01.PrintRune(rune(table[y][x]))
				z01.PrintRune(32)
			} else {
				z01.PrintRune(rune(table[y][x]))
			}
		}
		z01.PrintRune(10)
	}
}

// Input function confirms if the input is valid,
func checkInput(args []string) bool {
	if len(args) != 9 {
		fmt.Println("Error") // Input length is out of range
		return false
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error") //  Input length is out of range
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			//check if it equals char / or 0 and aslo
			//check if it less than . or greater than 9
			// if value == 47 || value == 48 {
			// 	fmt.Println("Error") // Input is not correct
			// 	return false
			// } else
			if value < 49 && value > 57 {
				fmt.Println("Error") // Input is not correct
				return false
			}
		}
	}
	return true
}

// Fill function fills the sudoku array of runes
// with values from the input string array.
func fillTable(table [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// CountDots function returns true if there are still
// empty cells, and returns false if there are none.
func isDots(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// IsValid function iterates over numbers
// that can be potentially placed in empty cells
// and checks if they are valid candidates
// (satisfy each of 3 conditions
// for number placement in sudoku game)
func isValid(table *[9][9]rune, x int, y int, z rune) bool {
	// check double int
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}

	//square check
	a := x / 3
	b := y / 3

	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}

// Solve function is a backtracking algorithm
// that uses recursion to check potential solutions
// starting from 1.
func isSolve(table *[9][9]rune) bool {
	//checks for dots in the table, if it doesnt find any, it stops
	if !isDots(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if isValid(table, x, y, z) {
						table[y][x] = z
						if isSolve(table) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}
