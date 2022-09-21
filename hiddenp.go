package main

import (
	"os"

	"github.com/01-edu/z01"
)

// ALLOWED IMPORTS: os.Args ,	z01.PrintRune

// Write a program named hiddenp that takes two strings and that, if the first string is hidden in the second one, displays 1 followed by a newline, otherwise it displays 0 followed by a newline.

// Let s1 and s2 be strings. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

// the first string is os.Args[1] and the second string is os.Args[2]

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	if os.Args[1] == "" {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	if len(str1) > len(str2) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	// check if characters in str1 are present in str2 in the same order
	success := false
	indexMatchedinS2 := -1
	numMatched := 0
	for j, c := range str1 {
		for i, c2 := range str2 {
			if c == c2 && i > indexMatchedinS2 && j == len(str1)-1 && numMatched == len(str1)-1 {
				success = true
				break
			} else if c == c2 && i > indexMatchedinS2 {
				indexMatchedinS2 = i
				numMatched++
				break
			}
		}
	}
	if success {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}

// EXAMPLES:
// $ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
// 1$
// $ go run . "abc" "2altrb53c.sse" | cat -e
// 1$
// $ go run . "abc" "btarc" | cat -e
// 0$
// $ go run . "DD" "DABC" | cat -e
// 0$
// $ go run .
// $
