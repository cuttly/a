package main

// Write a function that:

// converts an int value to a string using the specified base in the argument
// and that returns this string
// The base is expressed as an int, from 2 to 16. The characters comprising the base are the digits from 0 to 9, followed by uppercase letters from A to F.

// For example, the base 4 would be the equivalent of "0123" and the base 16 would be the equivalent of "0123456789ABCDEF".

// If the value is negative, the resulting string has to be preceded by a minus sign -.

// Only valid inputs will be tested.

// probable PREDICTIONS for tests/audits
// 1. The MAX-INT is never crossed (it appears from test cases)
// 2. Negative numbers between 100000 and -100000 are TESTED
// 3. The program is tested with different bases between 2 and 16

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}
	var sArr []rune
	minus := 1
	if value < 0 {
		sArr = append(sArr, '-')
		minus = -1
	}
	for value != 0 {
		sArr = append(sArr, rune(minus*(value%base))+'0')
		value /= base
	}
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] > '9' {
			sArr[i] = rune(sArr[i] - '9' + 'A' - 1)
		}
	}
	rArr := make([]rune, len(sArr))
	if minus == -1 {
		rArr[0] = '-'
		for i, j := 0, 1; i < len(sArr)-1; i++ {
			rArr[j] = sArr[len(sArr)-1-i]
			j++
		}
	} else {
		for i := 0; i < len(sArr); i++ {
			rArr[i] = sArr[len(sArr)-1-i]
		}
	}
	return string(rArr)
}

// func testItoaBase(value, base int) string {
// 	if base < 2 || base > 16 {
// 		return ""
// 	}
// 	return strings.ToUpper(strconv.FormatInt(int64(value), base))
// }

// func main() {
// 	fmt.Println(ItoaBase(-424575, 15))
// 	fmt.Println(testItoaBase(-424575, 15))
// }
