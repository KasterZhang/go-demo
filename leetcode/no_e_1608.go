package leetcode

import "sort"

//DONE
func specialArray1608(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := 0; i < len(nums); i++ {
		l := i + 1
		if nums[i] >= l {
			if i == len(nums)-1 {
				return l
			} else if nums[i+1] < l {
				return l
			}
		}
	}
	return -1
}
