package main

import "fmt"

func main() {
	PrintBits(2)
}

// If 2 is passed to the function PrintBits, it will print "00000010".
func PrintBits(octe byte) {
	for i := 7; i >= 0; i-- {
		if octe&(1<<i) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
