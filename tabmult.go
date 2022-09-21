package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: z01.PrintRune ,	os.* ,	strconv.Atoi

// Write a program that displays a number's multiplication table.

// The number is supplied as an os.Args[1] argument.

// The parameter will always be a strictly positive number that fits in an int, and said number times 9 will also fit in an int.

// Examples:
// Command: go run tabmult.go 9
// Output:
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	for i := 1; i <= 9; i++ {
		// only use of z01.PrintRune allowed, fmt is not allowed
		z01.PrintRune(rune(i + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		printInti(num)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printInti(num * i)
		z01.PrintRune('\n')
	}
}

func printInti(num int) {
	// divide by 10 and get the remainder and print it
	for num > 0 {
		rem := num % 10
		defer z01.PrintRune(rune(rem + '0'))
		num = num / 10
	}
}
