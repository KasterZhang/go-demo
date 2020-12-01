package leetcode

//DONE
func searchRange34(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start, end := 0, len(nums)-1
	findS, findE := false, false
	mid := (start + end) / 2
	for !findS || !findE {
		if end < start {
			return []int{-1, -1}
		}
		if nums[mid] > target {
			end = mid - 1
			mid = (start + end + 1) / 2
		} else if nums[mid] < target {
			start = mid + 1
			mid = (start + end) / 2
		} else if !findS {
			if mid == 0 || nums[mid-1] != target {
				findS = true
				start = mid
			} else {
				mid = (start + mid) / 2
			}
		} else {
			if mid+1 == len(nums) || nums[mid+1] != target {
				findE = true
				end = mid
			} else {
				mid = (mid + end + 1) / 2
			}
		}
	}

	return []int{start, end}
}
