package main

import "fmt"

func main() {
	displayalrevm()
}

func displayalrevm() {
	// Write a program that displays the alphabet in reverse,
	// with even letters in uppercase, and odd letters in lowercase,
	// followed by a newline('\n').
	for i := 122; i >= 97; i-- {
		if i%2 == 0 {
			fmt.Printf("%c", i)
		} else {
			fmt.Printf("%c", i-32)
		}
	}
}
