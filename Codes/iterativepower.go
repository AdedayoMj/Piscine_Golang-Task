package piscine

func IterativePower(nb int, power int) int {
	output := 1
	if power < 0 {
		output = 0
	} else {
		for power != 0 {
			output *= nb
			power -= 1
		}
	}
	return output
}
