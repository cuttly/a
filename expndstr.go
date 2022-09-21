package main

import "fmt"

func expstr() {
	str := " lol lol lol "
	a := []rune(str)
	b := len(a) - 1
	c := ""
	var d string
	var e []string
	x := 0
	y := 0
	if a[0] == ' ' {
		for i := 0; i < len(a); i++ {
			if a[i] != ' ' {
				x = i
				break
			}
		}
	}
	if a[b] == ' ' {
		for i := b; i > 0; i-- {
			if a[i] != ' ' {
				y = i
				break
			}
		}
	}
	for j := x; j <= y; j++ {
		c += string(a[j])
	}
	for _, r := range c {
		if r != ' ' {
			d += string(r)
		} else if r == ' ' {
			if d != "" {
				e = append(e, d)
				d = ""
			}
		}
	}
	if d != "" {
		e = append(e, d)
	}
	z := len(e)
	for g := 0; g < z; g++ {
		fmt.Print(e[g])
		if g != z-1 {
			fmt.Print(" ", " ", " ")
		}
	}
}
