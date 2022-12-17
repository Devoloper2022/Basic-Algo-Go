package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ans := os.Args

	for i := len(ans) - 1; i > 0; i-- {
		for j := 0; j < len(ans[i]); j++ {
			z01.PrintRune(rune(ans[i][j]))
		}
		z01.PrintRune('\n')
	}
}
