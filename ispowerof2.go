package main

func IsPowerOf2(n int) bool {
	if n%2 == 0 {
		for i := 2; true; i *= 2 {
			if i == n {
				return true
			} else {
				if i > n || i < 0 {
					return false
				}
			}
		}
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(IsPowerOf2(8))
// 	fmt.Println(IsPowerOf2(9))
// 	fmt.Println(IsPowerOf2(6))
// 	fmt.Println(IsPowerOf2(16))
// }
