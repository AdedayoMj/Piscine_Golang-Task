package main

import (
	"os"
)

func isSigne(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func Atoi(s string) (int, int) {
	myString := []rune(s)
	const Zero = 48
	res := 0
	length := len([]rune(s))

	if length == 0 {
		return 0, 0
	}

	var sign bool
	if myString[0] == 45 {
		myString[0] = 48
		sign = true
	} else {
		if myString[0] == 43 {
			myString[0] = 48
		}
	}
	for _, v := range myString {
		if v < 48 || v > 57 {
			return 0, 1
		}
	}
	for _, v := range myString {
		a := int(v - Zero)
		res = res*10 + a
	}
	if sign {
		res = res * -1
	}

	return res, 0
}

func main() {
	signe := []string{"+", "*", "-", "/", "%"}
	args := os.Args[1:]

	if len(args) != 3 {
	} else {
		if isSigne(args[1], signe) {
			nb1, err := Atoi(args[0])
			nb2, err2 := Atoi(args[2])

			// if AlphaCount(args[0]) || AlphaCount(args[2]) {
			output := 0
			checkItoa := itoa(nb1)
			if len(checkItoa) < 5 {
				if err == 0 && err2 == 0 {
					switch args[1] {
					case "+":
						output = nb1 + nb2
						os.Stderr.WriteString(itoa(output))
					case "-":
						output := nb1 - nb2
						os.Stderr.WriteString(itoa(output))
					case "/":
						if nb2 == 0 {
							os.Stderr.WriteString("No division by 0")
						} else {
							output = nb1 / nb2
							os.Stderr.WriteString(itoa(output))
						}
					case "%":
						if nb2 == 0 {
							os.Stderr.WriteString("No modulo by 0")
						} else {
							output = nb1 % nb2
							os.Stderr.WriteString(itoa(output))
						}
					case "*":
						output = nb1 * nb2
						os.Stderr.WriteString(itoa(output))
					}

					os.Stderr.WriteString(string('\n'))
				}
			}
		}
	}
}

func IntToString(num int) string {
	if num == 0 {
		return "0"
	} else if num == 1 {
		return "1"
	} else if num == 2 {
		return "2"
	} else if num == 3 {
		return "3"
	} else if num == 4 {
		return "4"
	} else if num == 5 {
		return "5"
	} else if num == 6 {
		return "6"
	} else if num == 7 {
		return "7"
	} else if num == 8 {
		return "8"
	} else {
		return "9"
	}
}

func IntSignToString(num int) string {
	if num == 0 {
		return "0"
	} else if num == -1 {
		return "1"
	} else if num == -2 {
		return "2"
	} else if num == -3 {
		return "3"
	} else if num == -4 {
		return "4"
	} else if num == -5 {
		return "5"
	} else if num == -6 {
		return "6"
	} else if num == -7 {
		return "7"
	} else if num == -8 {
		return "8"
	} else {
		return "9"
	}
}

func itoa(num int) string {
	var result string

	if num == 0 {
		return "0"
	}
	for num > 0 {
		result = IntToString(num%10) + result
		num /= 10
	}

	if num < 0 {
		for num < 0 {
			result = IntSignToString(num%10) + result
			num /= 10
		}
		result = "-" + result
	}

	return result
}
