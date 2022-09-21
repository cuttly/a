package main

import (
	"fmt"
)

func main() {
	displayalpham()
}

func displayalpham() {
	for i := 97; i <= 122; i++ {
		if i%2 == 0 {
			fmt.Printf("%c", i-32)
		} else {
			fmt.Printf("%c", i)
		}
	}
	fmt.Println()
}
