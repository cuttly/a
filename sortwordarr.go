package main

// ALLOWED PACKAGE: none

func SortWordArr(array []string) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

// TESTING BELOW
// func main() {
// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	SortWordArr(result)

// 	fmt.Println(result)
// }
