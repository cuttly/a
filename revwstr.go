package main

// Write a program that takes a string as a parameter, and prints its words in reverse, followed by a newline.

// A word is a sequence of alphanumerical characters.

// If the number of arguments is different from 1, the program will display nothing.

// In the parameters that are going to be tested, there will not be any extra spaces.
// (meaning that there will not be additional spaces at the beginning or at the end of the string and that words will always be separated by exactly one space).

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: os.Args ,	z01.PrintRune

func main() {
	if len(os.Args) != 2 {
		return
	}
	// if os.Args[1] == "" {
	// 	z01.PrintRune('\n')
	// 	return
	// }
	args := []rune(os.Args[1])

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}
	// in args, remove all spaces that are at the beginning or at the end of the []rune
	for args[0] == ' ' {
		args = args[1:]
	}
	for args[len(args)-1] == ' ' {
		args = args[:len(args)-1]
	}
	// if there are multiple space in between words, make them into one space
	for i := 0; i < len(args)-1; i++ {
		if args[i] == ' ' && args[i+1] == ' ' {
			args = append(args[:i], args[i+1:]...)
			i--
		}
	}

	// A word is a sequence of alphanumerical characters.
	allWords := [][]rune{}
	thisWord := []rune{}

	for i, ch := range args {
		if ch != ' ' {
			thisWord = append(thisWord, ch)
		}
		if ch == ' ' {
			allWords = append(allWords, thisWord)
			thisWord = []rune{}
		}
		if i == len(args)-1 {
			allWords = append(allWords, thisWord)
			thisWord = []rune{}
		}
	}
	// reverse the order of allWords
	// for i := 0; i < len(allWords)/2; i++ {
	// 	allWords[i], allWords[len(allWords)-1-i] = allWords[len(allWords)-1-i], allWords[i]
	// }
	// print out the array of array of runes
	if len(allWords) != 0 {
		for i := len(allWords) - 1; i >= 0; i-- {
			printRuneArr1(allWords[i])
			if i != 0 {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printRuneArr1(rstr []rune) {
	for _, ch := range rstr {
		z01.PrintRune(ch)
	}
}

// func printStr1(str string) {
// 	rstr := []rune(str)
// 	for _, ch := range rstr {
// 		z01.PrintRune(ch)
// 	}
// }

// $ go run . "the time of contempt precedes that of indifference"
// indifference of that precedes contempt of time the
// $ go run . "abcdefghijklm"
// abcdefghijklm
// $ go run . "he stared at the mountain"
// mountain the at stared he
// $ go run . "" | cat-e
// $
// $
