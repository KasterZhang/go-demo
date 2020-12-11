package leetcode

//DONE 线性筛 埃式筛 积性函数
func countPrimes(n int) int {
	result := 0
	nums := make([]int, n)
	for i, nn := range nums {
		if i >= 2 && nn != 1 {
			result++
			for j := i; j*i < n; j++ {
				nums[j*i] = 1
			}
		}
	}

	return result
}
