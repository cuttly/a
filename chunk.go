package main

import "fmt"

// Write a function called Chunk that receives as parameters a slice, slice []int, and an number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

// If the size is 0 it should print \n

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}
	var result [][]int
	for i := 0; i < len(slice); i += size {
		if i+size > len(slice) {
			result = append(result, slice[i:])
		} else {
			result = append(result, slice[i:i+size])
		}
	}
	fmt.Println(result)
}

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }

// EXPECTED OUTPUT => DO NOT UNCOMMENT BELOW THIS, THIS IS JUST FOR VERIFICATION
// []

// [[0 1 2] [3 4 5] [6 7]]
// [[0 1 2 3 4] [5 6 7]]
// [[0 1 2 3] [4 5 6 7]]
