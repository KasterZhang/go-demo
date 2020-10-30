package leetcode

func fourSum(nums []int, target int) [][]int {
	var qs func(arr []int) []int
	qs = func(arr []int) []int {
		if len(arr) == 0 {
			return []int{}
		}
		base, left, right := arr[0], []int{}, []int{}
		for _, v := range arr[1:] {
			if v <= base {
				left = append(left, v)
			}
			if v > base {
				right = append(right, v)
			}
		}
		left = qs(left)
		right = qs(right)
		return append(append(left, base), right...)
	}
	nums = qs(nums)
	result := [][4]int{}
	for i0, v0 := range nums {
		for i1, v1 := range nums[1+i0:] {
			for i2, v2 := range nums[2+i1+i0:] {
				for _, v3 := range nums[3+i2+i1+i0:] {
					if v0+v1+v2+v3 == target {
						first := true
						temp := [4]int{v0, v1, v2, v3}
						for _, v := range result {
							if v == temp {
								first = false
							}
						}
						if first {
							result = append(result, temp)
						}
					}
				}
			}
		}
	}
	temp := [][]int{}
	for i := range result {
		temp = append(temp, result[i][:])
	}

	return temp
}
