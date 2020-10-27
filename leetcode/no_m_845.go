package leetcode

//DONE
func longestMountain845(A []int) int {
	result := 0
	if len(A) < 3 {
		return result
	}
	pre := A[0]
	result++
	claim := true
	for i, val := range A[1:] {
		if claim {
			if val > pre {
				pre = val
				result++
			} else if val == pre {
				//等于的点位重新计数
				pre = val
				result = 1
			} else {
				if result > 1 {
					claim = false
					pre = val
					result++
				} else {
					pre = val
					result = 1
				}
			}
		} else {
			//下降阶段
			if val < pre {
				pre = val
				result++
			} else {
				next := longestMountain845(A[i:])
				if next > result {
					return next
				}
				return result

			}
		}
	}
	if !claim {
		return result
	}
	return 0
}
