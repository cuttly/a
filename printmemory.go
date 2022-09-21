package main

// Write a function that takes (arr [10]byte), and displays the memory as in the example.

// After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

// The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.

import (
	"unicode"

	"github.com/01-edu/z01"
)

// // ALLOWED IMPORTS: z01.PrintRune ,	unicode.IsGraphic

func PrintMemory(arr [10]byte) {
	// EXPECTED OUTPUT:=>
	// $ go run . | cat -e
	// 68 65 6c 6c$
	// 6f 10 15 2a$
	// 00 00$
	// hello..*..$
	// $

	// print all bytes as it is // ONLY FOR DEBUG
	// for _, ch := range arr {
	// 	fmt.Println(int(ch))
	// }
	// fmt.Println("done")

	// get hex value of each byte
	hexVal := []string{}

	for _, val := range arr {
		if val == 0 {
			hexVal = append(hexVal, "00")
		} else {
			hexVal = append(hexVal, itoaBase(int(val), 16))
		}
	}
	// populate final output rune array
	finalOut := []rune{}
	for i := 0; i < len(hexVal); i++ {
		// print hex value
		for j := 0; j < len(hexVal[i]); j++ {
			if unicode.IsGraphic(rune(hexVal[i][j])) {
				finalOut = append(finalOut, rune(hexVal[i][j]))
			}
		}
		// print space
		if ((i+1)%4 == 0 && i != 0) || i == len(hexVal)-1 {
			finalOut = append(finalOut, '\n')
		} else {
			finalOut = append(finalOut, ' ')
		}
	}

	for i := 0; i < 10; i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			finalOut = append(finalOut, rune(arr[i]))
		} else {
			finalOut = append(finalOut, '.')
		}
	}
	finalOut = append(finalOut, '\n')

	// if you need to add a 0 in front of the hex value, you can use this
	// wherever there is a rune between '0' and '9', and not preceded or postceded by another rune between '0' and '9', insert a '0' rune in front of such a rune in the rune arr
	for i := 0; i < len(finalOut); i++ {
		if finalOut[i] >= '0' && finalOut[i] <= '9' {
			if i != 0 && !(finalOut[i-1] >= '0' && finalOut[i-1] <= '9') && !(finalOut[i+1] >= '0' && finalOut[i+1] <= '9') && !(finalOut[i+1] >= 'a' && finalOut[i+1] <= 'z') {
				finalOut = append(finalOut[:i], append([]rune{'0'}, finalOut[i:]...)...)
			}
		}
	}

	printRuneArr(finalOut)
}

func printRuneArr(r []rune) {
	for _, ch := range r {
		z01.PrintRune(ch)
	}
}

// EXAMPLES:
// Usage
// Here is a possible program to test your function :

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// EXPECTED OUTPUT:=>
// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $

func itoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}
	var sArr []rune
	minus := 1
	if value < 0 {
		sArr = append(sArr, '-')
		minus = -1
	}
	for value != 0 {
		sArr = append(sArr, rune(minus*(value%base))+'0')
		value /= base
	}
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] > '9' {
			sArr[i] = rune(sArr[i] - '9' + 'a' - 1) // to capitalize alphabets, we need to add 'A' - 1
		}
	}
	rArr := make([]rune, len(sArr))
	if minus == -1 {
		rArr[0] = '-'
		for i, j := 0, 1; i < len(sArr)-1; i++ {
			rArr[j] = sArr[len(sArr)-1-i]
			j++
		}
	} else {
		for i := 0; i < len(sArr); i++ {
			rArr[i] = sArr[len(sArr)-1-i]
		}
	}
	return string(rArr)
}

//// ANOTHER SOLUTION IF REQUIRED

// func PrintMemory(a [10]int) {
// 	r := []rune{}
// 	for i, nbr := range a {
// 		fmt.Printf("%.2x", nbr)
// 		if ((i+1)%4 == 0 && i != 0) || i == len(a)-1 {
// 			z01.PrintRune('\n')
// 		} else {
// 			z01.PrintRune(' ')
// 		}

// 		if nbr >= 33 && nbr <= 126 {
// 			r = append(r, rune(nbr))
// 		} else {
// 			r = append(r, '.')
// 		}
// 	}
// 	for i := 0; i < 10-len(a); i++ {
// 		r = append(r, '.')
// 	}
// 	printRuneArr(r)
// }

// func printRuneArr(r []rune) {
// 	for _, ch := range r {
// 		z01.PrintRune(ch)
// 	}
// }

// // EXAMPLES:
// // Usage
// // Here is a possible program to test your function :

// func main() {
// 	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
// 	PrintMemory(arr)
// }

// // And its output :

// // go run printmemory.go | cat -e
// // 68 65 6c 6c$
// // 6f 10 15 2a$
// // 00 00$
// // hello..*..$
// // student@ubuntu:~/piscine/test$
