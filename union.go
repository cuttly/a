package main

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: os.Args , z01.PrintRune

func main() {
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		os.Exit(0)
	}
	args := os.Args[1:]
	str := ""
	if len(args) < 2 {
		z01.PrintRune('\n')
	} else {
		str = args[0] + args[1]
		str = removeDuplicates(str)
	}
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func removeDuplicates(str string) string {
	var result string
	for _, char := range str {
		if !contains(result, char) {
			result += string(char)
		}
	}
	return result
}

func contains(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}
