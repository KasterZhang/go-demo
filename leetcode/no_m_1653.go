package leetcode

//DONE
func minimumDeletions1653(s string) int {

	dp := make([]int, len(s))
	if len(dp) < 2 {
		return 0
	}
	for i, n := range s {
		temp := 0
		if n != 'a' {
			temp++
		}
		if i == 0 {
			dp[i] = temp
		} else {
			dp[i] = dp[i-1] + temp
		}
	}

	ndp := make([]int, len(s))
	for i, n := range s {
		temp := 0
		if n != 'b' {
			temp++
		}
		if i == 0 {
			ndp[i] = temp
		} else {
			if ndp[i-1] < dp[i] {
				ndp[i] = ndp[i-1] + temp
			} else {
				ndp[i] = dp[i-1] + temp
			}
		}
	}

	if dp[len(dp)-1] < ndp[len(ndp)-1] {
		return dp[len(dp)-1]
	}
	return ndp[len(ndp)-1]
}
