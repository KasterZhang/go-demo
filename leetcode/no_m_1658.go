package leetcode

//DONE
func minOperations1658(nums []int, x int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, 0
	lIndex, rIndex := -1, len(nums)-1
	result := -1
	for i, n := range nums {
		if left+n > x {
			break
		}
		lIndex = i
		left += n
		if left == x {
			result = i + 1
			break
		}
	}

	if lIndex == len(nums)-1 && result == -1 {
		return -1
	}

	for rIndex > 0 {
		if right+left+nums[rIndex] == x && (lIndex+1+len(nums)-1-rIndex < result || result == -1) {
			result = lIndex + len(nums) - rIndex + 1
		}
		if right+left+nums[rIndex] >= x {
			flag := false
			for lIndex >= 0 {
				left -= nums[lIndex]
				lIndex--
				if right+left+nums[rIndex] <= x {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			break
		}
		right += nums[rIndex]
		rIndex--
	}

	return result
}
