package piscine

// Write a function Map that, for an int slice, applies a function of this type func(int) bool
// on each elements of that slice and returns a slice of all the return values.

func Map(f func(int) bool, arr []int) []bool {
	r := []bool{}
	for _, s := range arr {
		r = append(r, f(s))
	}

	return r
}
