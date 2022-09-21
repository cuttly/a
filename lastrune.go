package main

func LastRune(s string) rune {
	r := []rune(s)
	return r[len(r)-1]
}

// func main() {
// 	z01.PrintRune(LastRune("Hello!"))
// 	z01.PrintRune(LastRune("Salut!"))
// 	z01.PrintRune(LastRune("Ola!"))
// 	z01.PrintRune(LastRune("♥01"))
// 	z01.PrintRune('\n')
// }
