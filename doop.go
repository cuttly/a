package main

import (
	"os"
)

// ALLOWED IMPORTS: os.*

const (
	maxPos = 9223372036854775807
	maxNeg = -9223372036854775808
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	num1, err1 := Atoi(os.Args[1])
	num2, err2 := Atoi(os.Args[3])

	if err1 || err2 {
		return
	}
	switch os.Args[2] {
	case "+":
		sum(num1, num2)
	case "-":
		if num2 == maxNeg {
			sum(-(num2 + 1), num1+1)
		} else {
			sum(num1, -num2)
		}
	case "*":
		mult(num1, num2)
	case "/":
		div(num1, num2)
	case "%":
		mod(num1, num2)
	}
}

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, true
	}

	sArr := []rune(s)
	minus := 1
	n := 0

	if sArr[0] == '-' {
		minus = -1
		sArr[0] = '0'
	} else if sArr[0] == '+' {
		sArr[0] = '0'
	}

	for _, k := range sArr {
		if k < '0' || '9' < k {
			return 0, true
		}

		if 0 <= n && n <= maxPos/10 || n < 0 && maxNeg/10 <= n {
			n *= 10
			digit := minus * int(k-'0')
			if minus == 1 && maxPos-digit >= n || minus == -1 && maxNeg-digit <= n {
				n += digit
			} else {
				return 0, true
			}
		} else {
			return 0, true
		}
	}
	return n, false
}

func Println(s string) {
	os.Stdout.WriteString(s + "\n")
}

func PrintNum(n int) {
	if n == 0 {
		Println("0")
		return
	}
	var sArr []rune
	ans := ""
	minus := 1
	if n < 0 {
		ans += "-"
		minus = -1
	}
	for n != 0 {
		sArr = append(sArr, rune(minus*(n%10))+'0')
		n /= 10
	}
	for i := len(sArr) - 1; i >= 0; i-- {
		ans += string(sArr[i])
	}
	Println(ans)
}

func sum(a, b int) {
	if b < 0 && maxNeg-b <= a || b >= 0 && a <= maxPos-b {
		PrintNum(a + b)
	}
}

func mult(a, b int) {
	if a == 0 || b == 0 {
		PrintNum(0)
	} else {
		if a > 0 && b > 0 {
			if a <= maxPos/b {
				PrintNum(a * b)
			}
		} else if a < 0 && b < 0 {
			if -a <= maxPos/-b {
				PrintNum(a * b)
			}
		} else if a < 0 {
			if maxNeg/b <= a {
				PrintNum(a * b)
			}
		} else if b < 0 {
			if b == -1 && a <= maxPos || a <= maxNeg/b {
				PrintNum(a * b)
			}
		}
	}
}

func div(a, b int) {
	if b == 0 {
		Println("No division by 0")
	} else {
		PrintNum(a / b)
	}
}

func mod(a, b int) {
	if b == 0 {
		Println("No modulo by 0")
	} else {
		PrintNum(a % b)
	}
}
