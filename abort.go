package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	ans := bubbleSort(arr, len(arr))

	return ans[2]
}

func bubbleSort(a []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j+1]
				a[j+1] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}
