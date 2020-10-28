package leetcode

//给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
//如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
//1 <= arr.length <= 1000
// -1000 <= arr[i] <= 1000
func uniqueOccurrences1207S1(arr []int) bool {
	arr = qs1207(arr)
	m := make(map[int]bool)
	pre := arr[0]
	count := 1
	for _, val := range arr[1:] {
		if val == pre {
			count++
		} else {
			pre = val
			if _, ok := m[count]; ok {
				return false
			}
			m[count] = true
			count = 1
		}
	}
	if _, ok := m[count]; ok {
		return false
	}
	return true
}
func qs1207(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	base := arr[0]
	left, right := []int{}, []int{}
	for _, val := range arr[1:] {
		if val > base {
			right = append(right, val)
		} else {
			left = append(left, val)
		}
	}
	left = qs1207(left)
	right = qs1207(right)
	return append(append(left, base), right...)
}
