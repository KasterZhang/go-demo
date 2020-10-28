package leetcode

//GOOD
func sortedSquares977S1(A []int) []int {
	result := make([]int, len(A))
	p0, p1, index := 0, len(A)-1, len(A)-1
	s0, s1 := A[p0]*A[p0], A[p1]*A[p1]
	for p0 != p1 {
		if s0 >= s1 {
			result[index] = s0
			p0++
			s0 = A[p0] * A[p0]
		} else {
			result[index] = s1
			p1--
			s1 = A[p1] * A[p1]
		}
		index--
	}
	result[0] = s0
	return result
}

//BAD
func sortedSquares977S2(A []int) []int {
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
