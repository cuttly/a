package main

// ALLOWED IMPORTS: make

func Split(s, sep string) []string {
	var result []string
	var index int
	for i := 0; i < len(s); i++ {
		if len(sep) > len(s[i:]) {
			break
		}
		if s[i:i+len(sep)] == sep {
			result = append(result, s[index:i])
			index = i + len(sep)
		}
	}
	result = append(result, s[index:])
	return result
}

// TESTING
// func main() {
// 	str := "HelloHAhowHAareHAyou?"
// 	fmt.Println(Split(str, "HA"))
// }
