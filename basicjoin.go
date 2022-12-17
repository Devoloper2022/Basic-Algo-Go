package piscine

func BasicJoin(elems []string) string {
	ans := ""

	for i := 0; i < len(elems); i++ {
		if i == len(elems)-1 {
			ans += elems[i]
		} else {
			ans += elems[i]
		}
	}

	return ans
}
