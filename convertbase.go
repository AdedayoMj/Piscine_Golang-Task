package piscine

// Write a function that prints the words of a string slice
// that will be returned by a function SplitWhiteSpaces.
// Each word is on a single line (each word ends with a \n).

func StringLength(str string) int {
	newStr := []rune(str)
	count := 0
	for range newStr {
		count++
	}
	return count
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n1 := StringLength(baseFrom)
	n2 := StringLength(baseTo)
	len := StringLength(nbr)
	num := 0
	inc := 1
	for i := len - 1; i >= 0; i-- {
		a := rune(nbr[i])
		for j, b := range baseFrom {
			if a == b {
				num = num + inc*j
			}
		}
		inc = inc * n1
	}
	res := ""
	inc = n2
	for i := 0; i < 100; i++ {
		if num < n2 {
			res = string(baseTo[num]) + res
			break
		}
		a := num % n2
		res = string(baseTo[a]) + res
		num = num / n2
	}

	return res
}
