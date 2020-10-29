package leetcode

//DONE
func twoSumS1(nums []int, target int) []int {
	for i, n := range nums {
		if n < target {
			for j, m := range nums {
				if n+m == target && i != j {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}
