package piscine

func isAlpha(alpha rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}

func AlphaCount(str string) int {
	count := 0
	for _, s := range str {
		if isAlpha(s) {
			count++
		}
	}
	return count
}
