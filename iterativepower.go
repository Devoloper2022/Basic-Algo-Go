package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	ans := nb
	for i := 0; i < power-1; i++ {
		ans = ans * nb
	}

	return ans
}
