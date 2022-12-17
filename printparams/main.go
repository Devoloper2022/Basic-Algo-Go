package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ans := os.Args

	for i := 1; i < len(ans); i++ {
		for j := 0; j < len(ans[i]); j++ {
			z01.PrintRune(rune(ans[i][j]))
		}
		z01.PrintRune('\n')
	}
}
