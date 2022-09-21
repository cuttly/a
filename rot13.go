package main

// Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

// 'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program displays nothing.

import "os"

// ALLOWED IMPORTS: z01.PrintRune ,	os.*

func main() {
	if len(os.Args) != 2 {
		return
	}
	if len(os.Args) == 2 {
		for _, r := range []rune(os.Args[1]) {
			if r >= 'a' && r <= 'z' {
				if r >= 'a'+13 {
					r -= 13
				} else {
					r += 13
				}
			} else if r >= 'A' && r <= 'Z' {
				if r >= 'A'+13 {
					r -= 13
				} else {
					r += 13
				}
			}
			os.Stdout.WriteString(string(r)) // comment out if illegal import, and make the below line with z01 run
			// z01.PrintRune(r)
		}
		os.Stdout.WriteString(string("\n")) // comment out if illegal import, and make the below line with z01 run
		// z01.PrintRune('\n')
	}
}

// FOR EXAMPLE:

// $ go run . "abc"
// nop
// $ go run . "hello there"
// uryyb gurer
// $ go run . "HELLO, HELP"
// URYYB, URYC
// $ go run .
// $
