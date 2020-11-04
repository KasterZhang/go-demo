package leetcode

// 当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

// 若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
// 或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
// 也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

// 返回 A 的最大湍流子数组的长度。
// 输入：[9,4,2,10,7,8,8,1,9]
// 输出：5
// 解释：(A[1] > A[2] < A[3] > A[4] < A[5])
//DONE
func maxTurbulenceSize978(A []int) int {
	dp := make([]int, len(A))
	if len(A) < 2 {
		return len(A)
	}
	dp[0] = 1
	max := 0
	if A[0] == A[1] {
		dp[1] = 1
		max = 1
	} else {
		dp[1] = 2
		max = 2
	}
	for i, n := range A[2:] {
		i = i + 2
		if (n > A[i-1] && A[i-1] < A[i-2]) || (n < A[i-1] && A[i-1] > A[i-2]) {
			dp[i] = dp[i-1] + 1
		} else if n == A[i-1] {
			dp[i] = 1
		} else {
			dp[i] = 2
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
