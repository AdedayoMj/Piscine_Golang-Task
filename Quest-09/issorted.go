package piscine

// Write a function IsSorted that returns true, if the slice of int is sorted, and that returns false otherwise.
// The function passed in parameter returns a positive int if a (the first argument) is superior to b
// (the second argument), it returns 0 if they are equal and it returns a negative int otherwise.
// To do your testing you have to write your own f function.

func IsSorted(f func(a, b int) int, tab []int) bool {
	if len(tab) > 1 {
		if f(tab[0], tab[1]) >= 0 {
			for i := 0; i < len(tab)-1; i++ {
				if f(tab[i], tab[i+1]) < 0 {
					return false
				}
			}
		}
		if f(tab[0], tab[1]) <= 0 {
			for i := 0; i < len(tab)-1; i++ {
				if f(tab[i], tab[i+1]) > 0 {
					return false
				}
			}
		}
	}

	return true
}
