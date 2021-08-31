package piscine

func Sqrt(nb int) int {
	if nb == 0 || nb == 1 {
		return nb
	} else {
		i := 1
		result := 1
		for result <= nb {
			i++
			result = i * i
		}
		if ((i - 1) * (i - 1)) != nb {
			return 0
		} else {
			return i - 1
		}
	}
}
