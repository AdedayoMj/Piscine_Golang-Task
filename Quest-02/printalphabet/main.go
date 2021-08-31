package main

import (
	"fmt"
)

func main() {
	// for ch := 97; ch <= 122; ch++ {
	// 	z01.PrintRune(ch)
	// }
	// z01.PrintRune('a')
	// z01.PrintRune('\n')
	for ch := 97; ch <= 122; ch++ {
		fmt.Printf("%c   ", ch)
	}
	fmt.Println()
}
