package main

import (
	"fmt"
	"os"
)

// Write a program that takes an undefined number of strings in arguments.
// For each argument, the program prints on the standard output "OK" followed by a newline if the expression is correctly bracketed, otherwise it prints "Error" followed by a newline.

// Symbols considered as brackets are brackets ( and ), square brackets [ and ]and braces { and }. Every other symbols are simply ignored.

// An opening bracket must always be closed by the good closing bracket in the correct order. A string which do not contains any bracket is considered as a correctly bracketed string.

// If there is no arguments, the program must print only a newline.

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		return
	}
	pair1a, pair1b, pair2a, pair2b, pair3a, pair3b := 0, 0, 0, 0, 0, 0
	openedBrackets := []rune{}
	for _, v := range os.Args[1:] {
		for _, r := range v {
			switch r {
			case '(':
				openedBrackets = append(openedBrackets, '(')
				pair1a++
			case ')':
				if len(openedBrackets) >= 1 && openedBrackets[len(openedBrackets)-1] == '(' {
					// remove the last element
					openedBrackets = openedBrackets[:len(openedBrackets)-1]
				} else {
					fmt.Println("Error")
					return
				}
				pair1b++
				if pair1b > pair1a {
					fmt.Println("Error")
					return
				}
			case '[':
				openedBrackets = append(openedBrackets, '[')
				pair2a++
			case ']':
				if len(openedBrackets) >= 1 && openedBrackets[len(openedBrackets)-1] == '[' {

					openedBrackets = openedBrackets[:len(openedBrackets)-1]
				} else {
					fmt.Println("Error")
					return
				}
				pair2b++
				if pair2b > pair2a {
					fmt.Println("Error")
					return
				}
			case '{':
				openedBrackets = append(openedBrackets, '{')
				pair3a++
			case '}':
				if len(openedBrackets) >= 1 && openedBrackets[len(openedBrackets)-1] == '{' {
					openedBrackets = openedBrackets[:len(openedBrackets)-1]
				} else {
					fmt.Println("Error")
					return
				}
				pair3b++
				if pair3b > pair3a {
					fmt.Println("Error")
					return
				}
			}
		}
		if pair1a != pair1b || pair2a != pair2b || pair3a != pair3b {
			fmt.Println("Error")
			return
		}
		fmt.Println("OK")
		pair1a, pair1b, pair2a, pair2b, pair3a, pair3b = 0, 0, 0, 0, 0, 0
	}
}

// Examples:

// INPUT =>
// go run brackets.go '(johndoe)' | cat -e
// OUTPUT =>
// OK$

// INPUT =>
// go run brackets.go '([)]' | cat -e
// OUTPUT =>
// Error$

// INPUT =>
// go run brackets.go '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
// OUTPUT =>
// OK$
// OK$

// INPUT =>
// go run brackets.go ")()" | cat -e
// OUTPUT =>
// Error$

// INPUT =>
// go run brackets.go " fh)difh(d" | cat -e
// OUTPUT =>
// Error$

// INPUT =>
// go run brackets.go | cat -e
// OUTPUT =>
// $

// INPUT =>
// go run brackets.go "{2*[d - 3]/(12)}" | cat -e
// OUTPUT =>
// OK$
