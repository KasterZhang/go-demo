package leetcode

//DONE
func getSmallestString1663(n int, k int) string {
	result := make([]rune, n)
	for k != 0 {
		if k-(n-1) >= 26 {
			result[n-1] = 'z'
			k -= 26
		} else {
			result[n-1] = rune(k - n + 97)
			k = n - 1
		}
		n--
	}
	return string(result)
}
