package piscine

// Write a function that takes an int min and an int max as parameters.
// That function returns a slice of int with all the values between min and max.
// Min is included, and max is excluded.
// If min is superior or equal to max, a nil slice is returned.
// append is not allowed for this exercise.

func MakeRange(min, max int) []int {
	var emp []int
	if min >= max {
		return emp
	}
	res := make([]int, max-min)
	for i := min; i < max; i++ {
		res[i-min] = i
	}
	return res
}
