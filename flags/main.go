package main

import (
	"fmt"
	"os"
)

// Write a program that can have as arguments --insert (or -i), --order (or -o) and a string.
// This program should :
// Insert the string given to the --insert (or -i), in the string argument, if given.
// Order the string argument (in ASCII order) if given the flag --order (or -o).
// In case there are no arguments or if the flag --help (or -h) is given,
// it should print the options, as shown in the example.

type Config struct {
	Ins   string
	Ord   bool
	Help  bool
	Input string
}

func ParsArgs() Config {
	var config Config
	if len(os.Args) < 2 {
		config.Help = true
		return config
	}
	for _, v := range os.Args[1:] {
		if len(v) >= 9 && v[0:9] == "--insert=" {
			config.Ins = v[9:]
		} else if len(v) >= 3 && v[0:3] == "-i=" {
			config.Ins = v[3:]
		} else if v == "--order" || v == "-o" {
			config.Ord = true
		} else if v == "--help" || v == "-h" {
			config.Help = true
		} else {
			if config.Input == "" {
				config.Input = v
			}
		}
	}
	return config
}

func main() {
	c := ParsArgs()
	if c.Help {
		helpMessage := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`
		fmt.Println(helpMessage)
		return
	}
	res := c.Input + c.Ins
	if c.Ord {
		res = Sort(res)
	}
	fmt.Println(res)
}

func Sort(str string) string {
	s := []byte(str)
	for i := range s {
		for j := i; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return string(s)
}
