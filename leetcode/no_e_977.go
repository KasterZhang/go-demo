package leetcode

func sortedSquares977(A []int) []int {
	result, b, c := make([]int, 0, len(A)), []int{}, []int{}
	for _, v := range A {
		if v <= 0 {
			b = append(b, v*v)
		} else {
			c = append(c, v*v)
		}
	}
	headB, headC := -1, -1

	for len(result) != len(A) {
		if len(b) != 0 {
			headB = b[len(b)-1]
		} else {
			headB = -1
		}
		if len(c) != 0 {
			headC = c[0]
		} else {
			headC = -1
		}
		if (headB > headC && headC != -1) || (headB == -1 && headC != -1) {
			result = append(result, headC)
			if len(c) != 0 {
				c = c[1:]
			}
		}
		if headC >= headB && headB != -1 || (headB != -1 && headC == -1) {
			result = append(result, headB)
			if len(b) != 0 {
				b = b[:len(b)-1]
			}
		}
	}
	return result
}
