package piscine

func ToLower(s string) string {
	ans := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			ans[i] = s[i] + 32
		} else {
			ans[i] = s[i]
		}
	}
	return string(ans)
}
