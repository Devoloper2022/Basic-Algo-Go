package main

import (
	"fmt"
	"piscine"
)

func main() {
	// fmt.Println(piscine.Atoi("12345"))
	// fmt.Println(piscine.Atoi("0000000012345"))
	// fmt.Println(piscine.Atoi("012 345"))
	// fmt.Println(piscine.Atoi("Hello World!"))
	// fmt.Println(piscine.Atoi("+1234"))
	// fmt.Println(piscine.Atoi("-1234"))
	// fmt.Println(piscine.Atoi("++1234"))
	// fmt.Println(piscine.Atoi("--1234"))

	s := []int{5, 4, 3, 2, 1, 0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
