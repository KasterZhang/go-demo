package leetcode

import (
	"fmt"
)

//DONE
func sortColors75S1(nums []int) {
	count, p0, p2 := 0, 0, len(nums)-1
	if p0 == p2 {
		return
	}
	for i, n := range nums {
		for {
			if count >= len(nums) || i > p2 {
				return
			}
			if n == 2 {
				if i != p2 {
					n = nums[p2]
					nums[p2] = 2
					nums[i] = n
					p2--
				} else {
					p2 = i - 1
					break
				}
				continue
			}
			if n == 0 {
				if i != p0 {
					n = nums[p0]
					nums[p0] = 0
					nums[i] = n
					p0++
				} else {
					p0 = i + 1
					break
				}
				continue
			}
			break

		}
		count++
	}
}

//DONE append替换首元素并位移不算原地操作 且复杂度不为n
func sortColors75S2(nums []int) {
	step := 0
	index := 0
	for step != len(nums) {
		if nums[index] == 2 {
			// nums[:index] = nums[:index+1]
			// nums[len(nums)-1] = 2
			nums = append(append(nums[:index], nums[index+1:]...), 2)
		} else if nums[index] == 0 {
			// nums[1:index] = nums[0:index]
			// nums[0] = 0
			nums = append([]int{0}, append(nums[:index], nums[index+1:]...)...)
			index++
		} else {
			index++
		}
		step++
	}
	fmt.Println(nums)
}
