package leetcode

func moveZeroes283S1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			j := i
			for j < len(nums)-1 && nums[j+1] != 0 {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
				j++
			}
		}
	}
}

func moveZeroes283S2(nums []int) {
	zero := 0
	for i, n := range nums {
		if n != 0 {
			nums[zero] = nums[i]
			zero++
		}
	}

	for zero < len(nums) {
		nums[zero] = 0
		zero++
	}
}
