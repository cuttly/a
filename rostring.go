package main

// Write a program that takes a string and displays this string after rotating it one word to the left.

// Thus, the first word becomes the last, and others stay in the same order.

// A word is a sequence of alphanumerical characters.

// Words will be separated by only one space in the output.

// If the number of arguments is different from 1, the program displays a newline.

// TESTING

// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
// $
// $

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: os.Args ,	z01.PrintRune

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
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

	if len(allWords) == 1 {
		printRuneArr2(allWords[0])
		z01.PrintRune('\n')
		return
	}

	// reverse the order of allWords
	// for i := 0; i < len(allWords)/2; i++ {
	// 	allWords[i], allWords[len(allWords)-1-i] = allWords[len(allWords)-1-i], allWords[i]
	// }
	// print out the array of array of runes

	// the last rune array in allWords
	lastWord := allWords[0]
	// remove the first rune array in allWords
	allWords = allWords[1:]

	// printRuneArr2(lastWord)
	// fmt.Println()

	if len(allWords) != 0 {
		for i := 0; i < len(allWords); i++ {
			printRuneArr2(allWords[i])
			z01.PrintRune(' ')
		}
		printRuneArr2(lastWord)
	}
	z01.PrintRune('\n')
}

func printRuneArr2(rstr []rune) {
	for _, ch := range rstr {
		z01.PrintRune(ch)
	}
}
