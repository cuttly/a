package main

import (
	"github.com/01-edu/z01"
)

// Write a program that displays all digits in descending order, followed by a newline('\n').
// input => go run countdown.go
// output => 9876543210

// allowed imports : z01.PrintRune

func main() {
	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
