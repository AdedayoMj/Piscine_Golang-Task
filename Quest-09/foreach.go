package piscine

// Write a function ForEach that, for an int slice, applies a function on each elements of that slice.

func ForEach(f func(int), arr []int) {
	for _, i := range arr {
		f(i)
	}
}
