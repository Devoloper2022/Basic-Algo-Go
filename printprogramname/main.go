package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ans := []rune(os.Args[0])

	for i := 2; i < len(ans); i++ {
		z01.PrintRune(rune(ans[i]))
	}
	z01.PrintRune('\n')
}
