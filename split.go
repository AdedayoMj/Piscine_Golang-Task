package piscine

// Write a function that separates the words of a string and
// puts them in a string slice.
// The separators are the characters of the separator string given in parameter.

func StringLen(str string) int {
	newStr := []rune(str)
	count := 0
	for range newStr {
		count++
	}
	return count
}

func Split(str, charset string) []string {
	str = str + charset
	len := StringLen(str)
	sepLen := StringLen(charset)
	wCount := 0
	for i := 0; i < len-sepLen+1; i++ {
		tmp := ""
		for j := 0; j < sepLen; j++ {
			tmp = tmp + string(str[i+j])
		}
		if tmp == charset {
			wCount++
		}
	}

	res := make([]string, wCount)
	wCount = 0
	pos := 0
	for i := 0; i < len-sepLen+1; i++ {
		tmp := ""
		word := ""
		pl := 0
		for j := 0; j < sepLen; j++ {
			tmp = tmp + string(str[i+j])
			pl = j
		}
		if tmp == charset {
			word = str[pos:i]
			res[wCount] = word
			wCount++
			pos = i + pl + 1
		}
	}
	return res
}
