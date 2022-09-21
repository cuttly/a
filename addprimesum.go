package main

import (
	"os"
)

// IMPORTS ALLOWED: os.* ,	z01.PrintRune

// Write a program that takes a positive integer as argument
// and displays the sum of all prime numbers inferior or equal to it followed by a newline.

// If the number of arguments is not 1, or if the argument is not a positive number,
// the program displays 0 followed by a newline.

// the number is os.Args[1]

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("0\n")
		// if os.Stdout not allowed, replace with below lines:
		// z01.PrintRune('0')
		// z01.PrintRune('\n')
		return
	}
	if os.Args[1] == "" {
		os.Stdout.WriteString("0\n")
		// if os.Stdout not allowed, replace with below lines:
		// z01.PrintRune('0')
		// z01.PrintRune('\n')
		return
	}
	num := myAtoi(os.Args[1])
	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	printThisNum(sum)
}

func printThisNum(num int) {
	if num == 0 {
		os.Stdout.WriteString("0\n")
		// if not allowed, replace with below lines:
		// z01.PrintRune('0')
		// z01.PrintRune('\n')
		return
	}
	var sArr []rune
	ans := ""
	minus := 1
	if num < 0 {
		ans += "-"
		minus = -1
	}
	for num != 0 {
		sArr = append(sArr, rune(minus*(num%10))+'0')
		num /= 10
	}
	for i := len(sArr) - 1; i >= 0; i-- {
		ans += string(sArr[i])
	}
	os.Stdout.WriteString(ans + "\n")
	// if os.Stdout not allowed, replace with below lines:
	// printThisString(ans)
	// z01.PrintRune('\n')
	// printThisString(ans)
}

// Uncomment this func if os.Stdout is not allowed
// func printThisString(str string) {
// 	rstr := []rune(str)
// 	for _, ch := range rstr {
// 		z01.PrintRune(ch)
// 	}
// }

func myAtoi(str string) int {
	num := 0
	minus := 1
	for _, c := range str {
		if c < '0' || '9' < c {
			return 0
		}
		if c == '-' {
			minus = -1
		} else if c == '+' {
			minus = 1
		} else {
			num *= 10
			num += minus * int(c-'0')
		}
	}
	return num
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// ONLY TESTING BELOW
// $ go run . 5
// 10
// $ go run . 7
// 17
// $ go run . -2
// 0
// $ go run . 0
// 0
// $ go run .
// 0
// $ go run . 5 7
// 0
// $
