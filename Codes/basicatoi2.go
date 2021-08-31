package piscine

func BasicAtoi2(s string) int {
	myString := []rune(s)
	const Zero = 48
	res := 0
	for _, v := range myString {
		if v >= 48 && v <= 57 {
			a := int(v - Zero)
			res = res*10 + a
		}
		if v < 48 || v > 57 {
			return 0
		}

	}
	return res
}
