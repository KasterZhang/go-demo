package leetcode

//DONE
func uniquePaths62(m int, n int) int {
	dp := make([]int, m)
	for i := 0; i < n; i++ {
		for j := range dp {
			if j == 0 {
				dp[j] = 1
			} else {
				if i > 0 {
					dp[j] += dp[j-1]
				} else {
					dp[j] = dp[j-1]
				}
			}
		}
	}
	return dp[m-1]
}
