package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Write a program that behaves like a simplified tail command but which takes at least one file as argument.
// The only option to be handled is -c and will be used in all the tests as the first argument, with positive values.
// For this program the os package can be used.
// Handle the errors by returning a non-zero exit status but process all the files.
// If several files are given, print a newline and the file name between each one of them (see below).

func atoi(s string) (int, bool) {
	sum := 0
	minus := false
	for i, c := range s {
		if c == '-' && i == 0 {
			minus = true
		} else if c >= 48 && c <= 57 {
			l := toInt(c)
			sum = sum*10 + l
		} else {
			return 0, false
		}
	}
	if minus {
		sum = -sum
	}
	return sum, true
}

func toInt(a rune) int {
	s := 0
	for i := '0'; i < a; i++ {
		s++
	}
	return s
}
func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err == nil {
		return i
	} else {
		return 0
	}
}

func main() {
	args := os.Args[1:]
	if 0 == len(args) {
		fmt.Println(os.Args[0])
	} else {
		num, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println(err.Error())
		} else {
			if 0 == len(os.Args[2:]) {
				fmt.Println("File name missing")
			} else {
				for _, res := range os.Args[2:] {
					ztail(res, num)
				}
			}
		}
	}
}

func ztail(s string, numBytes int) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", data[len(data)-numBytes:])
}

// func numberOfBytes(args []string) (int, []string) {
// 	n := len(args)
// 	nbytes := 0
// 	var files []string
// 	for i, v := range args {
// 		var err bool

// 		_, err = atoi(v)
// 		if v == "-c" {
// 			if i >= n-1 {
// 				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
// 				os.Exit(1)
// 			}
// 			arg := args[i+1]

// 			nbytes, err = atoi(arg)

// 			if err != true {
// 				fmt.Printf("tail: invalid number of bytes: %s\n", arg)
// 				os.Exit(1)
// 			}
// 			continue
// 		}

// 		if err != true {
// 			files = append(files, v)
// 		}

// 	}
// 	return nbytes, files
// }

// func fileSize(fi *os.File) int64 {
// 	fil, err := fi.Stat()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return 0
// 	}
// 	return fil.Size()
// }

// func main() {
// 	n := len(os.Args)
// 	if n < 4 {
// 		fmt.Println("Not enough arguments")
// 		// os.Exit(1)
// 	}

// 	nbytes, files := numberOfBytes(os.Args[1:])
// 	if nbytes > 0 {
// 		printName := len(files) > 1
// 		for j, f := range files {
// 			fi, err := os.Open(f)
// 			if err != nil {
// 				fmt.Printf("open '%s' for reading: No such file or directory\n", f)
// 				// os.Exit(0)
// 			}
// 			if printName {
// 				fmt.Printf("==> %s <==\n", f)
// 			}
// 			read := make([]byte, int(nbytes))
// 			_, er := fi.ReadAt(read, fileSize(fi)-int64(nbytes))
// 			if er != nil {
// 				// fmt.Println(er.Error())
// 			}

// 			for _, c := range read {
// 				z01.PrintRune(rune(c))
// 			}

// 			if j < len(files)-1 {
// 				z01.PrintRune('\n')
// 			}

// 			fi.Close()
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
