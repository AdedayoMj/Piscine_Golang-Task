package piscine

func IsNumeric(str string) bool {
	for _, a := range str {
		if !(a >= '0' && a <= '9') {
			return false
		}
	}
	return true
}

// func isDigit(digit rune) bool {
// 	for a := '0'; a <= '9'; a++ {
// 		if digit == a {
// 			return true
// 		}
// 	}
// 	return false
// }

// func IsNumeric(str string) bool {
// 	runeArray := []rune(str)
// 	runeCount := arrayCount(str)
// 	count := 0
// 	for _, char := range runeArray {
// 		if isDigit(char) {
// 			count++
// 		}
// 	}
// 	if count == runeCount {
// 		return true
// 	}
// 	return false
// }
