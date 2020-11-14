package leetcode

//DONE
func nextPermutation31(nums []int) {
	if len(nums) < 2 {
		return
	}
	pre := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if pre > nums[i] {
			tempI := i + 1
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					jt := nums[i]
					nums[i] = nums[j]
					nums[j] = jt
					break
				}
			}
			for j := len(nums) - 1; j > i; j-- {
				if tempI > j {
					break
				}
				t := nums[tempI]
				nums[tempI] = nums[j]
				nums[j] = t
				tempI++
			}
			break
		}
		if i == 0 {
			for j := len(nums) - 1; j > i; j-- {
				t := nums[i]
				nums[i] = nums[j]
				nums[j] = t
				i++
			}
			break
		}
		pre = nums[i]

	}
	return
}
