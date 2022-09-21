package main

import (
	"fmt"
	"os"
)

// ALLOWED IMPORTS: os.Args	fmt.Printf

type roman struct {
	num        int
	romanDigit string
}

func main() {
	if len(os.Args) == 2 {
		// In case it gives errors in weird cases like He33ooLl, uncomment below
		// if allNumsInStr(os.Args[1]) {
		// 	fmt.Printf("ERROR: cannot convert to roman digit\n")
		// 	return
		// }

		nbr := atoi12(os.Args[1])
		if nbr >= 4000 || nbr == 0 {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
			return
		}
		patter := []roman{
			{num: 1000, romanDigit: "M"},
			{num: 900, romanDigit: "CM"},
			{num: 500, romanDigit: "D"},
			{num: 400, romanDigit: "CD"},
			{num: 100, romanDigit: "C"},
			{num: 90, romanDigit: "XC"},
			{num: 50, romanDigit: "L"},
			{num: 40, romanDigit: "XL"},
			{num: 10, romanDigit: "X"},
			{num: 9, romanDigit: "IX"},
			{num: 5, romanDigit: "V"},
			{num: 4, romanDigit: "IV"},
			{num: 1, romanDigit: "I"},
		}
		sumRoman, romandigit := print12(nbr, patter)
		// if sumRoman contains '+' as the first or last digit, remove it
		for sumRoman[0] == '+' {
			sumRoman = sumRoman[1:]
		}
		for sumRoman[len(sumRoman)-1] == '+' {
			sumRoman = sumRoman[:len(sumRoman)-1]
		}
		fmt.Printf(sumRoman + "\n" + romandigit + "\n")
	}
}

func print12(nbr int, patter []roman) (string, string) {
	var sumRomanDigit, result string
	for _, v := range patter {
		for nbr >= v.num {
			sumRomanDigit += v.romanDigit + "+"
			result += v.romanDigit
			nbr -= v.num
		}
	}
	sumRomanDigit = formatsum(sumRomanDigit, patter)
	return sumRomanDigit, result
}

func formatsum(a string, patter []roman) string {
	result2 := []string{}
	accumulator := ""
	for _, v := range a {
		if v == '+' && accumulator != "" {
			result2 = append(result2, accumulator)
			accumulator = ""
		} else {
			accumulator += string(v)
		}
	}

	for i, v := range result2 {
		if len(v) == 2 {
			result2[i] = string(result2[i][1]) + string(result2[i][0])
		}
	}
	// a = strings.Join(result2, "+")
	// reqrite above line without using string.join
	for i, v := range result2 {
		if i == 0 {
			a = v
		} else {
			a += "+" + v
		}
	}
	return a
}

func atoi12(s string) int {
	base := "0123456789"
	var i, j int
	var result int
	var valid bool
	var baseLen int
	var sLen int
	var sArray []rune
	var baseArray []rune

	baseLen = 0
	sLen = 0
	result = 0
	valid = true

	for i = range base {
		baseLen++
	}

	for i = range s {
		sLen++
	}

	if baseLen < 2 {
		valid = false
	}

	for i = range base {
		for j = range base {
			if i != j && base[i] == base[j] {
				valid = false
			}
		}
	}

	if base[0] == '+' || base[0] == '-' {
		valid = false
	}

	if !valid {
		return 0
	}

	sArray = []rune(s)
	baseArray = []rune(base)

	isNegative := false

	// check if first element of sArray is '-' or '+'
	if sArray[0] == '-' {
		isNegative = true
		sArray = sArray[1:]
	}

	if sArray[0] == '+' {
		isNegative = false
		sArray = sArray[1:]
	}

	for i = range sArray {
		for j = range baseArray {
			if sArray[i] == baseArray[j] {
				result = result*baseLen + j
			}
		}
	}

	if isNegative {
		result = -result
	}

	return result
}

func allNumsInStr(s string) bool {
	// check if all the runes in a string are between 0 and 9
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// TESTING BELOW

// $ go run . hello
// ERROR: cannot convert to roman digit
// $ go run . 123
// C+X+X+I+I+I
// CXXIII
// $ go run . 999
// (M-C)+(C-X)+(X-I)
// CMXCIX
// $ go run . 3999
// M+M+M+(M-C)+(C-X)+(X-I)
// MMMCMXCIX
// $ go run . 4000
// ERROR: cannot convert to roman digit
// $
