package piscine

func Capitalize(s string) string {
	ans := []byte(s)

	for i := 0; i < len(ans); i++ {
		if Isuppercase(ans[i]) {
			ans[i] = ans[i] + 32
		}
	}

	if Islowercase(ans[0]) {
		ans[0] = ans[0] - 32
	}

	for i := 1; i < len(ans)-1; i++ {
		if Isalphanumeric(ans[i-1]) && Islowercase(ans[i]) && (Isalphanumeric(ans[i+1]) || Isletter(ans[i+1])) {
			ans[i] = ans[i] - 32
		}
	}

	if Islowercase(ans[len(s)-1]) && Isalphanumeric(ans[len(s)-2]) {
		ans[len(s)-1] = ans[len(s)-1] - 32
	}

	return string(ans)
}

func Isalphanumeric(ch byte) bool {
	if Isnumeric(ch) || Islowercase(ch) || Isuppercase(ch) {
		return true
	}

	return false
}

func Isletter(ch byte) bool {
	if Isuppercase(ch) || Islowercase(ch) {
		return true
	}

	return false
}

func Isuppercase(ch byte) bool {
	if ch > 64 && ch < 91 {
		return true
	}
	return false
}

func Islowercase(ch byte) bool {
	if ch > 96 && ch < 123 {
		return true
	}
	return false
}

func Isnumeric(ch byte) bool {
	if ch > 47 && ch < 58 {
		return true
	}
	return false
}
