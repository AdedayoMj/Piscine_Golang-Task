package piscine

func Atoi(s string) int {
	myString := []rune(s)
	const Zero = 48
	res := 0
	length := len([]rune(s))

	if length == 0 {
		return 0
	}

	var sign bool
	if myString[0] == 45 {
		myString[0] = 48
		sign = true
	} else {
		if myString[0] == 43 {
			myString[0] = 48
		}
	}
	for _, v := range myString {
		if v < 48 || v > 57 {
			return 0
		}
	}

	for _, v := range myString {
		a := int(v - Zero)
		res = res*10 + a
	}
	if sign {
		res = res * -1
	}

	return res
}
