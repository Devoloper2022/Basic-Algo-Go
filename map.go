package piscine

func Map(f func(int) bool, a []int) []bool {
	ans := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		ans[i] = f(a[i])
	}

	return ans
}
