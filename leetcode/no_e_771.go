package leetcode

//DONE
func numJewelsInStones771(J string, S string) int {
	m := make(map[rune]bool)
	for _, j := range J {
		m[j] = true
	}
	result := 0

	for _, s := range S {
		if _, ok := m[s]; ok {
			result++
		}
	}

	return result
}
