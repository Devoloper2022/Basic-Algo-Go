package main

import (
	"fmt"
	"piscine"
)

func main() {
	// fmt.Println(piscine.AppendRange(5, 10))
	// fmt.Println(piscine.AppendRange(10, 5))

	// fmt.Println(piscine.MakeRange(5, 10))
	// fmt.Println(piscine.MakeRange(10, 5))

	// test := []string{"Hello", "how", "are", "you?"}
	// fmt.Println(piscine.ConcatParams(test))

	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
}
