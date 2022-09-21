package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		wdmatch(s1, s2)
	} else {
		fmt.Println()
	}
}

func wdmatch(s1, s2 string) {
	index := -1
	countMatch := 0
	for i := 0; i < len(s1); i++ {
		if len(s1) != countMatch {
			for j := 0; j < len(s2); j++ {
				if s1[i] == s2[j] && j > index {
					index = j
					countMatch++
					break
				}
			}
		}
	}
	if countMatch == len(s1) {
		fmt.Println(s1)
	} else {
		fmt.Println()
	}

}
