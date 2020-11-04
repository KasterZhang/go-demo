package leetcode

func validMountainArray941(A []int) bool {
	if len(A) < 3 {
		return false
	}
	claim := true
	for i, node := range A {
		if i > 0 {
			if claim {
				if node < A[i-1] && i > 1 {
					claim = false
				} else if node <= A[i-1] {
					return false
				}
			} else {
				if node >= A[i-1] {
					return false
				}
			}
		}
	}
	return !claim
}
