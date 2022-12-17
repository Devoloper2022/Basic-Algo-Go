package main

import (
	"fmt"
	"os"
)

var ans [9][9]int

func draw() {
	for _, row := range ans {
		for index, el := range row {
			fmt.Print(el)
			if index != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func canPut(x int, y int, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if ans[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if ans[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if ans[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	if y == 9 {
		return true
	}
	if ans[y][x] != 0 {
		return solve(next(x, y))
	} else {
		for i := range [9]int{} {
			v := i + 1
			if canPut(x, y, v) {
				ans[y][x] = v
				if solve(next(x, y)) {
					return true
				}
				ans[y][x] = 0
			}
		}
		return false
	}
}

func Change(s []string) {
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] > 48 && s[i][j] < 58 {
				ans[i-1][j] = int(s[i][j]) - 48
			} else if s[i][j] == '.' {
				ans[i-1][j] = 0
			}
		}
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

func main() {
	given := os.Args

	if IsValidInput(given) {

		Change(given)
		if solve(0, 0) {
			draw()
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
