package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ans1 := true
	ans2 := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ans1 = false
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			ans2 = false
		}
	}
	if ans1 || ans2 {
		return true
	} else {
		return false
	}
}
