package piscine

// Write an iterative function that returns the factorial of the int passed as parameter.

// Errors (non possible values or overflows) will return 0.
func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 || nb > 2147483647 {
		result = result * 0
	} else {
		for i := 1; i <= nb; i++ {
			result *= i
		}
	}
	return result
}
