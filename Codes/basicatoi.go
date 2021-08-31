package piscine

func BasicAtoi(s string) int {
	myString := []rune(s)
	const Zero = 48
	res := 0
	for _, v := range myString {
		a := int(v - Zero)
		res = res*10 + a
	}
	return res
}
