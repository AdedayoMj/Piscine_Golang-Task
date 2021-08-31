package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func Lent3(d []string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}

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
