package main

import (
	"fmt"
	"os"
)

// Write a program that takes an undefined number of arguments which could be considered as options and writes on the standard output a representation of those options as groups of bytes followed by a newline.

// An option is an argument that begins by a - and have multiple characters which could be :

// abcdefghijklmnopqrstuvwxyz
// All options are stocked in a single int and each options represents a bit of that int, and should be stocked like this :

// 00000000 00000000 00000000 00000000
// ******zy xwvutsrq ponmlkji hgfedcba
// Launching the program without arguments or with the -h flag activated must print all the valid options on the standard output, as shown on one of the following examples.

// A wrong option must print "Invalid Option" followed by a newline.

// Examples
// INPUT =>
// go run options.go | cat -e
// OUTPUT =>
// options: abcdefghijklmnopqrstuvwxyz$
// INPUT =>
// go run options.go -abc -ijk | cat -e
// OUTPUT =>
// 00000000 00000000 00000111 00000111$
// INPUT =>
// go run options.go -z | cat -e
// OUTPUT =>
// 00000010 00000000 00000000 00000000$
// INPUT =>
// go run options.go -abc -hijk | cat -e
// OUTPUT =>
// options: abcdefghijklmnopqrstuvwxyz$
// INPUT =>
// go run options.go -% | cat -e
// OUTPUT =>
// Invalid Option$

// THIS SOLUTION SURVIVES ON A PRAYER AS DETAILED TEST CASES WERE NOT FOUND

func main() {
	if len(os.Args) < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	var options [32]bool
	for _, v := range os.Args {
		if len(v) < 2 {
			fmt.Println("Invalid Option")
			return
		}
		if v[0] == '-' {
			if v[1] == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			// fill options
			for _, r := range v[1:] {
				if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) { // check if it's a letter
					fmt.Println("Invalid Option")
					return
				}
				options['z'-r+6] = true
			}
		}
	}
	for i, v := range options { // print options
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
		if v {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	fmt.Println()
}
