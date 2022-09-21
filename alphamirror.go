package main

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: z01.PrintRune	os.Args

// Write a program called AlphaMirror that takes a string and displays this string
// after replacing each alphabetical character by the opposite alphabetical
// character, followed by a newline.

// The final result will be followed by a newline ('\n').

// 'a' becomes 'z', 'Z' becomes 'A'
// 'd' becomes 'w', 'M' becomes 'N'
// and so on.

// If the number of arguments is not 1, the program displays nothing.

func alphamirror(s string) {
	rs := []rune(s)
	for _, c := range rs {
		if c >= 'a' && c <= 'z' {
			c = 'z' - c + 'a'
		} else if c >= 'A' && c <= 'Z' {
			c = 'Z' - c + 'A'
		}
		z01.PrintRune(c)
	}
	z01.PrintRune('\n') // new line here
}

func main() {
	if len(os.Args) != 2 { // If the number of arguments is not 1, the program displays nothing.
		return
	}
	alphamirror(os.Args[1])
}

// TESTING BELOW
// $ go run . "abc"
// zyx
// $ go run . "My horse is Amazing." | cat -e
// Nb slihv rh Znzarmt.$
// $ go run .
// $
