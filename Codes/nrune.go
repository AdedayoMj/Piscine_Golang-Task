package piscine

func NRune(s string, n int) rune {
	myRunes := []rune(s)
	length := len(myRunes)
	if n <= 0 || n > length {
		return 0
	}
	return myRunes[n-1]
	// var output rune
	// if n < 0 || n > len(s) {
	// } else {
	// 	for i := 1; i < len(s)-1; i++ {
	// 		output = rune(s[n-1])
	// 	}
	// }
	// return output
}
