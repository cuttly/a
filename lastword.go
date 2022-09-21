package main

// ALLOWED: z01.PrintRune	os.Args

// Write a program that takes a string and displays its last word, followed by a newline('\n').

// A word is a section of string delimited by spaces or by the start/end of the string.

// The output will be followed by a newline('\n').

// If the number of parameters is not 1, or if there are no words, the program displays a newline('\n').

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	str := os.Args[1]
	rstr := []rune(str) // HANDLES SPECIAL CASES OF FOREIGN LETTER WORDS, straight string loop can fail there

	// if all the characters are spaces, return
	allspaces := true
	for _, ch := range str {
		if ch != ' ' {
			allspaces = false
			break
		}
	}
	if allspaces {
		z01.PrintRune('\n')
		return
	}

	// A word is a section of string delimited by spaces or by the start/end of the string.
	allWords := [][]rune{}
	thisWord := []rune{}
	for _, ch := range rstr {
		if ch != ' ' {
			thisWord = append(thisWord, ch)
		}
		if (ch == ' ' || ch == rune(str[len(str)-1])) && len(thisWord) > 0 {
			allWords = append(allWords, thisWord)
			thisWord = []rune{}
		}
	}
	if len(allWords) != 0 {
		printStr(allWords[len(allWords)-1])
	}
}

func printStr(rstr []rune) {
	for _, ch := range rstr {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

// TESTING BELOW
// // INPUT:
// go run lastword.go "FOR PONY" | cat -e
// // OUTPUT:
// PONY$

// // INPUT:
// go run lastword.go "this        ...       is sparta, then again, maybe    not" | cat -e
// // OUTPUT:
// not$

// // INPUT:
// go run lastword.go "  " | cat -e
// // OUTPUT:
// $

// // INPUT:
// go run lastword.go "a" "b" | cat -e
// // OUTPUT:
// $

// // INPUT:
// go run lastword.go "  lorem,ipsum  " | cat -e
// // OUTPUT:
// lorem,ipsum$
