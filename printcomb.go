package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	// not effective solution
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i < j && j < k && i < '7' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('7')
	z01.PrintRune('8')
	z01.PrintRune('9')

	z01.PrintRune('\n')
}
