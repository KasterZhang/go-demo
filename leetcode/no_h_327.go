package leetcode

import "fmt"

func countRangeSum327S1(nums []int, lower int, upper int) int {
	result := 0
	for i := range nums {
		for j := 0; j <= i; j++ {
			sum := sum327(nums, j, i)
			if sum >= lower && sum <= upper {
				fmt.Println("find", j, i, sum)
				result++
			}
		}
	}
	return result
}

func sum327(nums []int, start, end int) int {
	result := 0
	for start-1 != end {
		result += nums[start]
		start++
	}
	return result
}

func countRangeSum327S2(nums []int, lower int, upper int) int {
	result := 0
	m := [len(nums)][len(nums)]int{}
	for i := range nums {
		for j := 0; j <= i; j++ {
			sum := sum327(nums, j, i)
			if sum >= lower && sum <= upper {
				fmt.Println("find", j, i, sum)
				result++
			}
		}
	}
	return result
}
