package main

import (
	"fmt"
	"os"
	"strconv"
)

// ALLOWED IMPORTS: fmt.* ,	os.Args ,	strconv.Atoi

// Write a program that takes two strings representing two strictly positive integers that fit in an int.

// Display their greatest common divisor followed by a newline (It is always a strictly positive integer).

// If the number of parameters is not 2, display a newline.

// All arguments tested will be valid positive int values.

// the numbers are os.Args[1] and os.Args[2]

func main() {
	if len(os.Args) < 3 {
		// fmt.Println()
		return
	}
	if len(os.Args) > 3 {
		fmt.Println()
		return
	}

	if os.Args[1] == "" || os.Args[2] == "" {
		// fmt.Println()
	}

	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])

	if num1 < 0 || num2 < 0 {
		fmt.Println()
		return
	}

	// find highest numbers that divides both num1 and num2, by decreasing numbers from num1 to 1
	for i := num1; i > 0; i-- {
		if num1%i == 0 && num2%i == 0 {
			fmt.Println(i)
			return
		}
	}
}

// EXAMPLES:
// $ go run . 42 10 | cat -e
// 2$
// $ go run . 42 12
// 6
// $ go run . 14 77
// 7
// $ go run . 17 3
// 1
// $ go run .
// $ go run . 50 12 4
// $
