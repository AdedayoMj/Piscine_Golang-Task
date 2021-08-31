package piscine

// Write a function Any that returns true, for a string slice :
// if, when that string slice is passed through an f function, at least one element returns true.

func Any(f func(string) bool, arr []string) bool {
	for _, i := range arr {
		if f(i) {
			return true
		}
	}
	return false
}
