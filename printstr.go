package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	l := len([]rune(s))
	for i := 0; i < l; i++ {
		z01.PrintRune(rune(s[i]))
	}
}
