package main

import (
	"fmt"
	"piscine"
)

func main() {
	// z01.PrintRune(piscine.FirstRune("Hello!"))
	// z01.PrintRune(piscine.FirstRune("Salut!"))
	// z01.PrintRune(piscine.FirstRune("Ola!"))
	// z01.PrintRune(piscine.FirstRune("♥01"))
	// z01.PrintRune('\n')

	// z01.PrintRune(piscine.LastRune("Hello!"))
	// z01.PrintRune(piscine.LastRune("Salut!"))
	// z01.PrintRune(piscine.LastRune("Ola!"))
	// z01.PrintRune('\n')

	// z01.PrintRune(piscine.NRune("Hello!", 3))
	// z01.PrintRune(piscine.NRune("Salut!", 2))
	// z01.PrintRune(piscine.NRune("Bye!", -1))
	// z01.PrintRune(piscine.NRune("Bye!", 5))
	// z01.PrintRune(piscine.NRune("Ola!", 4))
	// z01.PrintRune('\n')

	// fmt.Println(piscine.Compare("Hello!", "Hello!"))
	// fmt.Println(piscine.Compare("Salut!", "lut!"))
	// fmt.Println(piscine.Compare("Ola!", "Ol"))

	// s := "Hello 78 World!    4455 /"
	// nb := piscine.AlphaCount(s)
	// fmt.Println(nb)

	// fmt.Println(piscine.Concat("Hello!", " How are you?"))

	// fmt.Println(piscine.IsUpper("HELLO"))
	// fmt.Println(piscine.IsUpper("HELLO!"))

	// fmt.Println(piscine.IsLower("hello"))
	// fmt.Println(piscine.IsLower("hello!"))

	// fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	// fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	// fmt.Println(piscine.IsAlpha("What's this 4?"))
	// fmt.Println(piscine.IsAlpha("Whatsthis4"))

	// fmt.Println(piscine.IsNumeric("010203"))
	// fmt.Println(piscine.IsNumeric("01,02,03"))

	// fmt.Println(piscine.IsPrintable("Hello"))
	// fmt.Println(piscine.IsPrintable("Hello\n"))

	// fmt.Println(piscine.ToUpper("Hello! How are you?"))

	// fmt.Println(piscine.ToLower("Hello! How are you?"))

	// piscine.PrintNbrInOrder(321)
	// piscine.PrintNbrInOrder(0)
	// piscine.PrintNbrInOrder(321)

	// fmt.Println(piscine.TrimAtoi("12345"))
	// fmt.Println(piscine.TrimAtoi("str123ing45"))
	// fmt.Println(piscine.TrimAtoi("012 345"))
	// fmt.Println(piscine.TrimAtoi("Hello World!"))
	// fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	// fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	// fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	// fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))

	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(piscine.Capitalize("[-+],Pg`E~u~8"))
	fmt.Println(piscine.Capitalize("^WOAf3}KiWlQF"))
}
