package piscine

func IterativeFactorial(nb int) int {
	ans := 1
	if nb < 0 || nb > 39 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 0; i < nb; i++ {
			ans = ans * (i + 1)
		}
	}
	return ans
}
