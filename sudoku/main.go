package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ans := os.Args
	if IsValidInput(ans) {
		// Solve(ans)
		ShowResult(ans)
	}
}

func IsValidInput(s []string) bool {
	if len(s) != 10 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if len(s[i]) != 9 {
			return false
		}
		for j := 0; j < len(s[i]); j++ {
			if !(s[i][j] > 47 && s[i][j] < 58) && s[i][j] != 46 {
				return false
			}
		}
	}
	return true
}

// func solve(s []string, x int, y int) bool {
// 	if y == 9 {
// 		return true
// 	}
// 	if s[y][x] != 0 {
// 		x, y := MoveElement(x, y)
// 		return solve(s, x, y)
// 	} else {
// 		for i := range [9]int{} {
// 			v := i + 1
// 			if CheckDuplicate(s, x, y, v) {
// 				// d := s[y][x] - byte(v)
// 				s[y][x] = byte(v)
// 				x, y := MoveElement(x, y)
// 				if solve(s, x, y) {
// 					return true
// 				}
// 				s[y][x] = byte('.')
// 			}
// 		}
// 		return false
// 	}
// }

func MoveElement(x, y int) (int, int) {
	x = (x + 1) % 9
	if x == 0 {
		y += 1
	}
	return x, y
}

func VerticalCheck(s []string, x, y, v int) bool {
	for i := range [9]int{} {
		if s[i][x] == byte(v) {
			return true
		}
	}
	return false
}

func HorizontalCheck(s []string, x, y, v int) bool {
	for i := range [9]int{} {
		if s[y][i] == byte(v) {
			return true
		}
	}
	return false
}

func SquareCheck(s []string, x, y, v int) bool {
	spx, spy := int(x/3)*3, int(y/3)*3
	for chy := range [3]int{} {
		for chx := range [3]int{} {
			if s[spx+chy][spy+chx] == byte(v) {
				return true
			}
		}
	}
	return false
}

func ShowResult(s []string) {
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			z01.PrintRune(rune(s[i][j]))
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func CheckDuplicate(s []string, x, y, v int) bool {
	return !VerticalCheck(s, x, y, v) &&
		!HorizontalCheck(s, x, y, v) &&
		!SquareCheck(s, x, y, v)
}
