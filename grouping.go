package main

import (
	"fmt"
	"os"
)

// Allowed imports:=   os.Args	,  fmt.Printf

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("\n") // comment this out if test gives error i.e. they do not want new line when no arguments are passed
		return
	}
	if len(os.Args) > 3 {
		fmt.Printf("\n") // comment this out if test gives error i.e. they do not want new line when more than 2 strings supplied
		return
	}
	regular := []rune(os.Args[1])
	parent := []rune(os.Args[2])
	if len(regular) <= 0 || len(parent) <= 0 {
		fmt.Printf("\n") // comment this out if test gives error i.e. they do not want new line when empty string supplied
		return
	}
	// if regular does not contain ( or ) or ( and ) are not in the same order
	if !(regular[0] == '(' && regular[len(regular)-1] == ')') {
		fmt.Printf("\n") // comment this out if test gives error i.e. they do not want new line when regular is not in the correct format
		return
	}
	allChildren := [][]rune{}
	buffer := []rune{}
	// for each | in regular, add it to allChildren
	for i := 0; i < len(regular); i++ {
		if i != 0 && i != len(regular)-1 {
			buffer = append(buffer, regular[i])
		}
		if regular[i] == '|' {
			allChildren = append(allChildren, buffer[:len(buffer)-1])
			buffer = []rune{}
		}
		if i == len(regular)-1 {
			allChildren = append(allChildren, buffer)
			buffer = []rune{}
		}
	}
	result := [][]rune{}
	parentStrings := [][]rune{}
	// for each space in parent, add it to parentStrings
	for i := 0; i < len(parent); i++ {
		if i != 0 {
			buffer = append(buffer, parent[i])
		}
		if parent[i] == ' ' || i == len(parent)-1 || parent[i] == '\n' {
			parentStrings = append(parentStrings, buffer)
			buffer = []rune{}
		}
	}
	// for each string in allChildren, if it is in parent, add it to result
	for _, str := range parentStrings {
		// for each string in allChildren, if it is in parent, add it to result
		for _, thisToSearach := range allChildren {
			if isContained(thisToSearach, str) {
				result = append(result, str)
			}
		}
	}
	// for each string in result, remove leading and trailing ','
	for i := 0; i < len(result); i++ {
		for result[i][len(result[i])-1] == ',' || result[i][len(result[i])-1] == ' ' {
			result[i] = result[i][:len(result[i])-1]
		}
	}
	for i := 0; i < len(result); i++ {
		for result[i][0] == ',' || result[i][0] == ' ' {
			result[i] = result[i][1:]
		}
	}
	// if result is empty, print new line
	if len(result) == 0 {
		fmt.Printf("\n") // comment this out if test gives error i.e. they do not want new line when no match is found
	} else {
		// if result is not empty, print result
		for i, v := range result {
			fmt.Printf("%d: %v\n", i+1, string(v))
		}
	}
}

func isContained(child []rune, parent []rune) bool {
	for i := 0; i < len(parent)-len(child); i++ {
		if parent[i] == child[0] {
			// if all the elements of parent[i:i+len(child)] are the same as child
			for j := 0; j < len(child); j++ {
				if parent[i+j] != child[j] {
					break
				}
				if j == len(child)-1 {
					return true
				}
			}
		}
	}
	return false
}

// INPUT =>
// go run grouping.go "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
// OUTPUT =>
// 1: heavy
// 2: steady
// 3: heavy

// INPUT =>
// go run grouping.go "(e|n)" "I currently have 4 windows opened up… and I don’t know why."
// OUTPUT =>
// 1: currently
// 2: currently
// 3: have
// 4: windows
// 5: opened
// 6: opened
// 7: and
// 8: don’t
// 9: know

// INPUT =>
// go run grouping.go "(hi)" "He swore he just saw his sushi move."
// OUTPUT =>
// 1: his
// 2: sushi

// INPUT =>
// go run grouping.go "(s)" ""

// INPUT =>
// go run grouping.go "i" "Something in the air"
