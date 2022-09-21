package main

// Write a function AtoiBase that takes two arguments:

// s: a numeric string in a given base.
// base: a string representing all the different digits that can represent a numeric value.
// And return the integer value of s in the given base.

// If the base is not valid it returns 0.

// Validity rules for a base :

// A base must contain at least 2 characters.
// Each character of a base must be unique.
// A base should not contain + or - characters.
// String number must contain only elements that are in base.

// Only valid string numbers will be tested.

// The function does not have to manage negative numbers.

// ALLOWED IMPORTS : "z01.PrintRune"

func AtoiBase(s string, base string) int {
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

// TESTING

// func main() {
// 	fmt.Println(AtoiBase("125", "0123456789"))      // 125
// 	fmt.Println(AtoiBase("1111101", "01"))          // 125
// 	fmt.Println(AtoiBase("7D", "0123456789ABCDEF")) // 125
// 	fmt.Println(AtoiBase("uoi", "choumi"))          // 125
// 	fmt.Println(AtoiBase("bbbbbab", "-ab"))         // 0
// }
