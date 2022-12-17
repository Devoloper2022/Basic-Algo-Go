package piscine

func IsPrintable(s string) bool {
	ch := true
	for i := 0; i < len(s); i++ {
		if s[i] < 32 {
			return false
		}
	}
	return ch
}
