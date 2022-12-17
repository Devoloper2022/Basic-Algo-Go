package piscine

func ActiveBits(n int) int {
	i := n
	count := 0
	for i > 0 {

		if i%2 == 1 {
			count++
		}
		i /= 2
	}

	return count
}
