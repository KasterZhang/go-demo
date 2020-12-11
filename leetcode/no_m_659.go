package leetcode

//DONE
func isPossible659(nums []int) bool {
	l := len(nums)
	c, dict := make([]int, 1, l), make([]int, 1, l)
	c[0] = 1
	dict[0] = nums[0]
	for i, n := range nums[1:] {
		if nums[i] == n {
			c[len(c)-1]++
		} else {
			dict = append(dict, n)
			c = append(c, 1)
		}
	}

	for l != 0 {
		count := 0
		pre, preIndex := 0, 0
		for i, n := range dict {
			if c[i] != 0 {
				if pre+1 != n && count != 0 {
					break
				}
				if c[i] <= c[preIndex] && count != 0 {
					break
				}
				count++
				pre = n
				preIndex = i
				c[i]--
				l--
			}
		}
		if count < 3 {
			return false
		}
	}
	return true
}
