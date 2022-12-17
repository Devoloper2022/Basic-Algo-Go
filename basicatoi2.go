package piscine

func BasicAtoi2(s string) int {
	given := []rune(s)
	ans := 0
	for i := 0; i < len(s); i++ {
		if given[i] > 47 && given[i] < 58 {
			ans = ans * 10
			ans = ans + int(given[i]-48)
		} else {
			return 0
		}
	}

	return ans
}
