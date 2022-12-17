package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	ans := -1

	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			ans = i
			for j := 0; j < len(toFind); j++ {
				if len(s) <= i+j {
					ans = -1
					break
				}

				if s[i+j] != toFind[j] {
					ans = -1
					break
				}

			}
		}
		if ans != -1 {
			return ans
		}
	}

	return ans
}
