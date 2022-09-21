package main

import (
	"os"
	"strconv"
)

// ALLOWED IMPORTS: z01.PrintRune ,	os.* ,	strconv.*

// Write a program that takes a positive (or zero) number (os.Args[1]) expressed in base 10,
// and displays it in base 16 ( with lowercase letters) followed by a newline.

// If the number of parameters is not 1, the program displays a newline.

func main() {
	if len(os.Args) != 2 {
		return
	}
	if os.Args[1] == "" {
		os.Stdout.WriteString("\n")
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	// convert to base 16 hex such that "10" becomes "a", "255" becomes "ff", "5156454" becomes "4eae66" etc.
	var sArr []rune
	ans := ""
	minus := 1
	if num < 0 {
		ans += "-"
		minus = -1
	}
	for num != 0 {
		sArr = append(sArr, rune(minus*(num%16))+'0')
		num /= 16
	}
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] > '9' {
			ans += string(sArr[i] - '9' + 'a' - 1)
		} else {
			ans += string(sArr[i])
		}
	}
	if ans == "" {
		ans = "ERROR"
	}
	os.Stdout.WriteString(ans + "\n")
}

// EXAMPLES:
// $ go run . 10
// a
// $ go run . 255
// ff
// $ go run . 5156454
// 4eae66
// $ go run .
// $ go run . "123 132 1" | cat -e
// ERROR$
// $
