package main

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: z01.PrintRune	, os.Args

// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

// The display will be followed by a newline ('\n').

// If the number of arguments is different from 2, the program displays nothing.

// Examples of output:

// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]
	newstr := ""
	// thisStr := "len of str1: " + itoa1234(len(str1)) + "len of str2: " + itoa1234(len(str2))
	// for i := 0; i < len(thisStr); i++ {
	// 	z01.PrintRune([]rune(thisStr)[i])
	// }
	// z01.PrintRune('\n')
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				newstr += string(str1[i])
				break
			}
		}
	}
	// remove all duplicates
	for i := 0; i < len(newstr); i++ {
		for j := i + 1; j < len(newstr); j++ {
			if newstr[i] == newstr[j] {
				newstr = newstr[:j] + newstr[j+1:]
				j--
			}
		}
	}
	for _, v := range newstr {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func itoa1234(n int) string {
	if n == 0 {
		return "0"
	}

	var result []rune
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		result = append(result, rune(n%10+'0'))
		n /= 10
	}

	if negative {
		result = append(result, '-')
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
