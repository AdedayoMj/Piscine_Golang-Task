package piscine

// Write a function that takes the arguments received in parameters
// and returns them as a string. The string is the result of all
// the arguments concatenated with a newline (\n) between

func ConcatParams(args []string) string {
	res := ""
	n := 0
	for range args {
		n++
	}
	for i := 0; i < n-1; i++ {
		res = res + args[i] + "\n"
	}
	res = res + args[n-1]
	return res
}
