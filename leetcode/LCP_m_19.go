package leetcode

//DONE
func minimumOperationsLCP19(leaves string) int {
	dp := make([][]int, 3)
	for i := range dp {
		line := make([]int, len(leaves))
		dp[i] = line
	}
	count := 0
	for i, n := range leaves[:len(leaves)-2] {
		if n != 'r' {
			count++
		}
		dp[0][i] = count
	}

	for i, n := range leaves[1 : len(leaves)-1] {
		temp := 0
		if n != 'y' {
			temp = 1
		}
		if i == 0 {
			dp[1][i+1] = dp[0][i] + temp
			continue
		}
		if dp[0][i] < dp[1][i] {
			dp[1][i+1] = dp[0][i] + temp
		} else {
			dp[1][i+1] = dp[1][i] + temp
		}
	}

	for i, n := range leaves[2:] {
		temp := 0
		if n != 'r' {
			temp = 1
		}
		if i == 0 {
			dp[2][i+2] = dp[1][i+1] + temp
		} else {
			if dp[1][i+1] < dp[2][i+1] {
				dp[2][i+2] = dp[1][i+1] + temp
			} else {
				dp[2][i+2] = dp[2][i+1] + temp
			}
		}
	}

	return dp[2][len(leaves)-1]
}
