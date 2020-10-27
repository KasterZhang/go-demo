package leetcode

import (
	"fmt"
)

func smallerNumbersThanCurrent1365S1(nums []int) []int {
	rowId := make([]int, len(nums), len(nums))
	result := make([]int, len(nums), len(nums))
	for i := range nums {
		rowId[i] = i
	}
	qsS1(nums, rowId, result, make(map[int]int), 0)
	return result
}

func qsS1(nums, location, result []int, replica map[int]int, offset int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	base := nums[0]
	left, right, leftLocation, rightLocation := []int{}, []int{}, []int{}, []int{}

	for i, v := range nums {
		if i != 0 {
			if v >= base {
				rightLocation = append(rightLocation, location[i])
				right = append(right, v)
			} else {
				leftLocation = append(leftLocation, location[i])
				left = append(left, v)
			}
		}
	}

	result[location[0]] = offset + len(left)
	if val, ok := replica[base]; ok {
		result[location[0]] -= val
		replica[base] = val + 1
	} else {
		replica[base] = 1
	}
	left = qsS1(left, leftLocation, result, replica, offset)
	right = qsS1(right, rightLocation, result, replica, offset+len(left)+1)
	return append(append(left, base), right...)
}

//TODO 可能不适合logN时间复杂度,付出的空间成本太高了,可以考虑下衡量空间和时间的成本
func smallerNumbersThanCurrent1365S2(nums []int) []int {
	rowId := make([]int, len(nums), len(nums))
	for i := range nums {
		rowId[i] = i
	}
	return qsS2(nums, rowId, 0)
}

func qsS2(nums []int, location []int, offset int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	base := nums[0]
	left, right, leftLocation, rightLocation := []int{}, []int{}, []int{}, []int{}
	fmt.Println("------------------")
	// fmt.Println("location",location)
	for i, v := range nums {
		if i != 0 {
			if v >= base {
				rightLocation = append(rightLocation, location[i])
				right = append(right, v)
			} else {
				leftLocation = append(leftLocation, location[i])
				left = append(left, v)
			}
		}
	}
	// fmt.Println("nums",nums)
	// fmt.Println("leftLocation",leftLocation)
	// fmt.Println("rightLocation",rightLocation)

	fmt.Println("value", base, "location", location[0], "len", offset+len(left))
	left = qsS2(left, leftLocation, offset)
	right = qsS2(right, rightLocation, offset+len(left)+1)
	return append(append(left, base), right...)
}
