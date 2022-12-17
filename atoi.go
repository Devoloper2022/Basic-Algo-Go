package piscine

func Atoi(s string) int {
	runes := []rune(s)

	result := 0

	if len(s) == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < len(runes)-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}

	if runes[0] == '-' {
		result *= -1
	}

	return result
}
