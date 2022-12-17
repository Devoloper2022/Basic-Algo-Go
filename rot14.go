package piscine

func Rot14(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			if s[i]+14 > 122 {
				ans += string(s[i] + 14 - 26)
			} else {
				ans += string(s[i] + 14)
			}
		} else if s[i] > 64 && s[i] < 91 {
			if s[i]+14 > 90 {
				ans += string(s[i] + 14 - 26)
			} else {
				ans += string(s[i] + 14)
			}
		} else {
			ans += string(s[i])
		}
	}
	return ans
}
