package piscine

func IsAlpha(s string) bool {
	ch := true
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 32 && s[i] < 48 || s[i] > 57 && s[i] < 65 || s[i] > 90 && s[i] < 97 || s[i] > 122 {
			return false
		}
	}
	return ch
}
