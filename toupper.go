package piscine

func ToUpper(s string) string {
	ans := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			ans[i] = s[i] - 32
		} else {
			ans[i] = s[i]
		}
	}
	return string(ans)
}
