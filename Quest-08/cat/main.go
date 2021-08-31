package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

// Write a program that behaves like a simplified cat command.
// The options do not have to be handled.
// If the program is called without arguments it should take the
// standard input (stdin) and print it back on the standard output (stdout).

func main() {
	lens := len(os.Args[1:])

	for i := 1; i <= lens; i++ {
		file, err := ioutil.ReadFile(os.Args[i])
		strFile := string(file)
		if err != nil {
			input := "ERROR: " + err.Error()
			os.Stderr.WriteString(input)
			z01.PrintRune('\n')
			os.Exit(1)
			z01.PrintRune('\n')

		} else {
			for _, v := range strFile {
				z01.PrintRune(v)
			}

			// z01.PrintRune('\n')
		}

	}
	if lens == 0 {
		for {
			text, _ := ioutil.ReadAll(os.Stdin)
			for _, v := range string(text) {
				z01.PrintRune(v)
			}
			os.Exit(0)
		}
	}
	os.Exit(0)
}
