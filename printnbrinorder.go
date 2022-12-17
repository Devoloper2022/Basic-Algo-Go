package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for i := 0; i < 10; i++ {
		for j := n; j > 0; j /= 10 {
			if j%10 == i {
				z01.PrintRune(rune(i + 48))
			}
		}
	}
}
