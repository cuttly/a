package main

// The function should have as parameters a slice of integers, a []int and a function f func(int, int) int. You should apply for each element of the slice the arithmetic function, saving it and printing.

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	println(result)
}

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }
