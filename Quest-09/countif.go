package piscine

// Write a function CountIf that returns, the number of
// elements of a string slice, for which the f function returns true.

func CountIf(f func(string) bool, tab []string) int {
	numberelements := 0
	for _, s := range tab {
		if f(s) == true {
			numberelements++
		}
	}
	return numberelements
}
