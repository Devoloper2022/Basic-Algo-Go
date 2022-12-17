package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	answer := 0
	array := []rune(s)
	tmp := 1
	if len(array) == 0 {
		return -1
	}
	if array[0] == '-' {
		tmp = -1
	}
	for i := len(s) - 1; i > 0; i-- {
		if array[i] < '0' || array[i] > '9' {
			return -1
		}
		if answer > 9223372036854775806-(int(array[i])-'0')*tmp {
			return -1
		}
		answer = answer + (int(array[i])-'0')*tmp
		tmp = tmp * 10
	}
	if array[0] != '+' && array[0] != '-' {
		if answer > 9223372036854775806-(int(array[0])-'0')*tmp {
			return -1
		}
		answer = answer + (int(array[0])-'0')*tmp
	}
	return answer
}

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		first := Atoi(args[0])
		second := Atoi(args[2])
		if first == -1 || second == -1 {
			return
		}
		if args[1] == "+" {
			fmt.Println(first + second)
		}
		if args[1] == "-" {
			fmt.Println(first - second)
		}
		if args[1] == "*" {
			fmt.Println(first * second)
		}
		if args[1] == "/" {
			if second == 0 {
				fmt.Println("No division by 0")
				return
			}
			fmt.Println(first / second)
		}
		if args[1] == "%" {
			if second == 0 {
				fmt.Println("No modulo by 0")
				return
			}
			fmt.Println(first % second)
		}
	}
}
