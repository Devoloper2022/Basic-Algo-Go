package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var ans []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		}
		if s[i] == ' ' && word != "" {
			ans = append(ans, word)
			word = ""
		}
	}

	if word != "" {
		ans = append(ans, word)
	}

	return ans
}
