package piscine

func BasicAtoi(s string) int {
	given := []rune(s)
	ans := 0
	for i := 0; i < len(s); i++ {
		ans = ans * 10
		ans = ans + int(given[i]-48)
	}

	return ans
}
