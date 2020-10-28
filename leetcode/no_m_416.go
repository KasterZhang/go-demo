package leetcode

import (
	"fmt"
)

//DONE 位运算 受长度限制
func canPartition416S1(nums []int) bool {
	sum, max, l := 0, 0, len(nums)
	if l < 2 {
		return false
	}
	for _, val := range nums {
		sum += val
		if val > max {
			max = val
		}
	}
	if sum&1 == 1 {
		return false
	}
	if max > sum>>1 {
		return false
	}

	var target byte = 1 << (sum / 2)
	var temp byte = target<<1 - 1
	fmt.Printf("%d %d %b %b\n", sum, sum/2, target, temp)
	dp := make([]byte, len(nums))
	for i := range dp {
		dp[i] = 1 << nums[i]
		if i > 0 {
			dp[i] = dp[i] | dp[i-1]
			dp[i] = dp[i] | dp[i-1]<<nums[i]&temp
		}
	}
	for i := range dp {
		fmt.Printf("%d %d  %b\n", nums[i], i, dp[i])
	}

	return dp[len(nums)-1]&target >= target
}

//DONE
func canPartition416S2(nums []int) bool {
	sum, max, l := 0, 0, len(nums)
	if l < 2 {
		return false
	}
	for _, val := range nums {
		sum += val
		if val > max {
			max = val
		}
	}
	if sum&1 == 1 {
		return false
	}
	target := sum >> 1
	if max > target {
		return false
	}
	//并不需要二维数组line即是preLine
	line := make([]bool, target+1)
	for i := l - 1; i > 0; i-- {
		if nums[i] == target {
			return true
		}
		if i > 0 {
			// newLine := make([]bool, target+1, target+1)
			// copy(newLine, line)
			for j, flag := range line {
				if flag && j+nums[i] <= target && !line[j+nums[i]] {
					line[j+nums[i]] = true
					if j+nums[i] == target {
						return true
					}
				}
			}
			// line = newLine
		}
		line[nums[i]] = true
	}
	return line[target]
}
