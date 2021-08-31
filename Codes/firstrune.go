package piscine

func FirstRune(s string) rune {
	var output rune
	for _, c := range s {
		output = c
		break
	}
	return output
}
