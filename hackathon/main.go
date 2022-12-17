package main

import (
	"fmt"
	"piscine"
)

const N = 6

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(a)
	fmt.Println(unmatch)
}
