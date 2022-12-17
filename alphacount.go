package piscine

func AlphaCount(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 || s[i] > 96 && s[i] < 123 {
			ans += 1
		}
	}
	return ans
}
