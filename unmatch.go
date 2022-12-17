package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		count := 1
		for j := 0; j < len(a); j++ {
			if a[j] == a[i] && i != j {
				count++
			}
		}
		if count%2 == 1 {
			return a[i]
		}

	}

	return -1
}
