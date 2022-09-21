package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ALLOWED IMPORTS: os.* ,	fmt.* ,	strings.Split ,	strconv.Atoi

func deleteExtraSpaces(a []string) (res []string) {
	for _, v := range a {
		if v != "" {
			res = append(res, v)
		}
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	var values []int
	op := strings.Split(os.Args[1], " ")
	op = deleteExtraSpaces(op)
	for _, v := range op {
		val, err := strconv.Atoi(v)

		if err == nil {
			values = append(values, val)
			continue
		}

		n := len(values)
		if n < 2 {
			fmt.Println("Error")
			return
		}

		switch v {
		case "+":
			values[n-2] += values[n-1]
			values = values[:n-1]
		case "-":
			values[n-2] -= values[n-1]
			values = values[:n-1]
		case "*":
			values[n-2] *= values[n-1]
			values = values[:n-1]
		case "/":
			values[n-2] /= values[n-1]
			values = values[:n-1]
		case "%":
			values[n-2] %= values[n-1]
			values = values[:n-1]
		default:
			fmt.Println("Error")
			return
		}
	}
	if len(values) == 1 {
		fmt.Println(values[0])
	} else {
		fmt.Println("Error")
	}
}

// // TESTING BELOW HERE

// // INPUT:
// go run rpncalc.go "1 2 * 3 * 4 +" | cat -e
// // OUTPUT:
// 10$

// // INPUT:
// go run rpncalc.go "1 2 3 4 +" | cat -e
// // OUTPUT:
// Error$

// // INPUT:
// go run rpncalc.go | cat -e
// // OUTPUT:
// Error$

// // INPUT:
// go run rpncalc.go "     1      3 * 2 -" | cat -e
// // OUTPUT:
// 1

// // INPUT:
// go run rpncalc.go "     1      3 * ksd 2 -" | cat -e
// // OUTPUT:
// Error$
