package piscine

func Join(strs []string, sep string) string {
	ans := ""
	ans += strs[0] + sep

	for i := 1; i < len(strs)-1; i++ {
		ans += strs[i] + sep
	}

	ans += strs[len(strs)-1]

	return ans
}
