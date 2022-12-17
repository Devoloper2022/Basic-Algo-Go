package piscine

func Compact(ptr *[]string) int {
	count := 0
	a := *ptr

	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			count++
		}
	}

	index := 0
	arr := make([]string, count)
	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			arr[index] = a[i]
			index++
		}
	}

	*ptr = arr
	return count
}
