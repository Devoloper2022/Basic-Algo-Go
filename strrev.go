package piscine

func StrRev(s string) string {
	l := len([]rune(s))

	given := []rune(s)
	ans := []rune(s)

	for i := 0; i < l; i++ {
		ans[i] = given[l-i-1]
	}
	return string(ans)
}
