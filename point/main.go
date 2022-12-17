package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func main() {
	points := &point{}

	setPoint(points)

	show(points.x, points.y)
}

func show(x, y int) {
	ans := "x = " + convert(x) + ", y = " + convert(y)

	for i := 0; i < len(ans); i++ {
		z01.PrintRune(rune(ans[i]))
	}
	z01.PrintRune('\n')
}

func convert(a int) string {
	temp := ""
	ans := ""
	for i := a; i > 0; i /= 10 {
		temp += string((i % 10) + 48)
	}

	for i := len(temp) - 1; i >= 0; i-- {
		ans += string(temp[i])
	}

	return ans
}
