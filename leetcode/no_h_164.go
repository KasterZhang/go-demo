package leetcode

import (
	"fmt"
)

//桶排
//DONE
func maximumGap164S1(nums []int) int {

	result, d, l, min, max := 0, 1, len(nums), 0, 0
	if l < 2 {
		return 0
	}
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	if (max-min)/(l-1) > 1 {
		d = (max - min) / (l - 1)
	}

	buckects := make([][]int, (max-min)/d+1)
	for i := range buckects {
		buckects[i] = []int{-1, -1}
	}
	for _, n := range nums {
		indext := n / d
		if n > buckects[indext][1] {
			buckects[indext][1] = n
		}
		if n < buckects[indext][0] || buckects[indext][0] == -1 {
			buckects[indext][0] = n
		}
	}
	fmt.Println(buckects, d, (max-min)/d+1)
	pre := buckects[0]
	for _, bucket := range buckects[1:] {
		if bucket[0] != -1 {
			if pre[0] == -1 {
				pre = bucket
			} else {
				if bucket[0]-pre[1] > result {
					result = bucket[0] - pre[1]
				}
				pre = bucket
			}
		}
	}
	return result
}

//基数排序
//DONE
func maximumGap164S2(nums []int) int {
	result, l, max := 0, len(nums), 0
	if l < 2 {
		return result
	}
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	ignore := 1
	for max/ignore != 0 {
		buckects := make([][]int, 10)
		for _, n := range nums {
			buckects[n/ignore%10] = append(buckects[n/ignore%10], n)
		}
		temp := make([]int, 0, l)
		for _, buckect := range buckects {
			temp = append(temp, buckect...)
		}
		nums = temp
		ignore *= 10
	}
	fmt.Println(nums, ignore)
	for i, n := range nums[1:] {
		if n-nums[i] > result {
			result = n - nums[i]
		}
	}
	return result
}

func maximumGap(nums []int) int {

	result, d, l, min, max := 0, 1, len(nums), 0, 0
	if l < 2 {
		return 0
	}
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	if (max-min)/(l-1) > 1 {
		d = (max - min) / (l - 1)
	}

	buckects := make([][]int, (max-min)/d+1)
	for i := range buckects {
		buckects[i] = []int{-1, -1}
	}
	for _, n := range nums {
		indext := n / d
		if n > buckects[indext][1] {
			buckects[indext][1] = n
		}
		if n < buckects[indext][0] || buckects[indext][0] == -1 {
			buckects[indext][0] = n
		}
	}

	pre := buckects[0]
	for _, bucket := range buckects[1:] {
		if bucket[0] != -1 {
			if pre[0] == -1 {
				pre = bucket
			} else {
				if bucket[0]-pre[1] > result {
					result = bucket[0] - pre[1]
				}
				pre = bucket
			}
		}
	}
	return result
}
