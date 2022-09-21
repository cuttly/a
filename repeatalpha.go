package main

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: z01.PrintRune	, os.*

// Write a program called repeat_alpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// The result must be followed by a newline ('\n').

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

// If the number of arguments is different from 1, the program displays nothing.

func main() {
	if len(os.Args) != 2 {
		return
	}
	rs := []rune(os.Args[1])
	for _, ch := range rs {
		if ch >= 'a' && ch <= 'z' {
			for i := 0; i < int(ch-'a'+1); i++ {
				z01.PrintRune(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			for i := 0; i < int(ch-'A'+1); i++ {
				z01.PrintRune(ch)
			}
		} else {
			z01.PrintRune(ch)
		}
	}

	z01.PrintRune('\n')
}

// Examples of usage:
// $ go run . abc | cat -e
// abbccc
// $ go run . Choumi. | cat -e
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $ go run . "abacadaba 01!" | cat -e
// abbacccaddddabba 01!$
// $ go run .
// $ go run . "" | cat -e
// $
// $
