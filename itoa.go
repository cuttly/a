package main

// Allowed imports:  len

// Write a function that simulates the behaviour of the Itoa function in Go.
// Itoa transforms a number represented as anint in a number represented as a string.

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result []rune
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		result = append(result, rune(n%10+'0'))
		n /= 10
	}

	if negative {
		result = append(result, '-')
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// func main() {
// 	os.Stdout.WriteString(Itoa(-1000) + "\n")
// }
