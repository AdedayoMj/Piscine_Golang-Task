package piscine

// Write a function that takes an int min and an int max as parameters.
// That function returns a slice of int with all the values between min and max.
// Min is included, and max is excluded.
// If min is superior or equal to max, a nil slice is returned.
// make is not allowed for this exercise.

func AppendRange(min, max int) []int {
	var res []int
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
