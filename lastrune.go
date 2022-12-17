package piscine

func LastRune(s string) rune {
	word := []rune(s)
	return rune(word[len(word)-1])
}
