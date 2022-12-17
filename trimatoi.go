package piscine

func TrimAtoi(s string) int {
	ans := 0
	ss := false
	chn := false

	for j := 0; j < len(s); j++ {
		if s[j] == 45 && chn == false {
			ss = true
		}

		if s[j] > 47 && s[j] < 58 {
			ans *= 10
			ans += int(s[j] - 48)
			chn = true
		}
	}

	if ss == true {
		ans *= -1
	}

	return ans
}
