package leetcode

//TODO
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			temp := nums[i-1]
			nums[i-1] = nums[i]
			nums[i] = temp
			break
		} else {
			temp := nums[i-1]
			nums[i-1] = nums[i]
			nums[i] = temp
		}
	}
}
