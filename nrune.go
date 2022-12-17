package piscine

func NRune(s string, n int) rune {
	word := []rune(s)

	a := len(word)
	if n <= 0 || a < n {
		return 0
	}

	return rune(word[n-1])
}
