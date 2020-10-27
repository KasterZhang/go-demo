package leetcode

//DONE
func partitionLabels763(S string) []int {
	d := make(map[rune]int)
	for i, r := range S {
		d[r] = i
	}
	result := []int{}
	target, pre := 0, 0

	for i, r := range S {
		last := d[r]
		if last > target+pre {
			target = last - pre
		}
		if target == i-pre {
			result = append(result, target+1)
			target = 0
			pre = i + 1
		}
	}
	return result
}
