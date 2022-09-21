package main

import (
	"fmt"
	"os"
	"strconv"
)

// ALLOWED IMPORTS :=> strconv.Atoi ,	os.Args ,	fmt.*

// Write a program that takes a positive int and displays its prime factors on the standard output, followed by a newline.
// If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

// Factors must be displayed in ascending order and separated by *,
// so that the expression in the output gives the right result.

// The input, when there is one, will always be valid.

// Example:
// when arg is 225225
// output is 3*3*5*5*7*11*13
// when arg is 8333325
// output is 3*3*5*5*7*11*13*37
// when arg is 9539
// output is 9539
// when arg is 804577
// output is 804577
// when arg is 42
// output is 2*3*7

// the bloew should give no outpiut:
// go run . a
// go run . 0
// go run . 1

func main() {
	if len(os.Args) != 2 {
		// os.Stdout.WriteString("\n")
		return
	}
	if os.Args[1] == "" {
		// os.Stdout.WriteString("\n")
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	if num == 0 {
		// os.Stdout.WriteString("0\n")
		return
	}
	if num < 0 {
		// os.Stdout.WriteString("\n")
		return
	}
	var factors []int
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			num /= i
			i--
		}
	}
	for i := 0; i < len(factors); i++ {
		if i == len(factors)-1 {
			fmt.Println(factors[i])
			// os.Stdout.WriteString(strconv.Itoa(factors[i]) + "\n")
		} else {
			fmt.Print(factors[i], "*")
			// os.Stdout.WriteString(strconv.Itoa(factors[i]) + "*")
		}
	}
}
