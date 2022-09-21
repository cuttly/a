package main

import (
	"github.com/01-edu/z01"
)

// The function should have as parameters a function, f func(int, int) int a slice of integers, slice []int and an int acc int.
// You should apply for each element of the slice the arithmetic function, saving and printing it.
// The function will be tested with our own functions Add, Sub, and Mul.

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, v := range a {
		result = f(result, v)
	}
	z01PrintNumLn(result)
}

func z01PrintNumLn(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	invertNum(&n)
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			z01.PrintRune(rune(n%10) + '0')
			n /= 10
		}
	}
	z01.PrintRune('\n')
}

func invertNum(n *int) { // invert number so that first digit becomes last digit and last digit becomes first digit
	var temp []int
	for *n > 0 {
		temp = append(temp, *n%10)
		*n /= 10
	}
	*n = 0
	for i := 0; i <= len(temp)-1; i++ {
		*n = *n*10 + temp[i]
	}
}

// TESTING BELOW

// func Add(a, b int) int {
// 	return a + b
// }

// func Sub(a, b int) int {
// 	return a - b
// }

// func Mul(a, b int) int {
// 	return a * b
// }

// func main() {
// 	table := []int{1, 2, 3}
// 	ac := 93
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// 	fmt.Println()

// 	table = []int{0}
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// }

// EXPECTED OUTPUT => DO NOT UNCOMMENT BELOW THIS, THIS IS JUST FOR VERIFICATION
// 99
// 558
// 87

// 93
// 0
// 93
