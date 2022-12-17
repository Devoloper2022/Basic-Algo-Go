package piscine

func CollatzCountdown(start int) int {
	count := 0
	if start < 1 {
		return -1
	}
	i := start

	for i > 0 {

		if i < 2 {
			break
		} else if i%2 == 1 {
			i = 3*i + 1
		} else {
			i /= 2
		}
		count++
	}

	return count
}
