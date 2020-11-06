package leetcode

//DONE
func lengthOfLIS300(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i, n := range nums[1:] {
		count := 1
		for j := len(nums[:i+1]) - 1; j >= 0; j-- {
			if n > nums[j] && count < dp[j]+1 {
				count = dp[j] + 1
				if dp[j] == max {
					break
				}
			}
		}
		if count > max {
			max = count
		}
		dp[i+1] = count
	}
	return max
}
