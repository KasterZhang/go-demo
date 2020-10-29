package leetcode

import (
	"fmt"
)

//TODO
func sortColors75S1(nums []int) {
	index0, index2 := 0, len(nums)-1
	for i, n := range nums {
		if n == 0 {
			temp := nums[index0]
			fmt.Println("temp", temp, "n", n, "i", i, "index0", index0)
			fmt.Println("nums", nums)
			nums[index0] = n
			nums[i] = temp
			fmt.Println("nums", nums)
			index0++
		}
		if n == 2 && index2 > i {
			temp := nums[index2]
			fmt.Println("temp", temp, "n", n, "i", i, "index2", index2)
			fmt.Println("nums", nums)
			nums[index2] = n
			nums[i] = temp
			fmt.Println("nums", nums)
			index2--
		}
		fmt.Println("-------------")
		fmt.Println(index0, index2)
	}
	fmt.Println(nums)
}

// index0, index1 := 0, 0
// for i, n := range nums {
// 	if n == 0 {
// 		temp := nums[index0]
// 		fmt.Println("temp", temp, "n", n, "i", i, "index0", index0)
// 		fmt.Println("nums", nums)
// 		nums[index0] = n
// 		nums[i] = temp
// 		fmt.Println("nums", nums)
// 		index0++
// 		index1++
// 	}
// 	if n == 1 {
// 		temp := nums[index1]
// 		fmt.Println("temp", temp, "n", n, "i", i, "index1", index1)
// 		fmt.Println("nums", nums)
// 		nums[index1] = n
// 		nums[i] = temp
// 		fmt.Println("nums", nums)
// 		index1++
// 	}
// 	fmt.Println("-------------")
// 	fmt.Println(index0, index1)
// }
// fmt.Println(nums)

//DONE append替换首元素并位移不算原地操作 且复杂度不为n
func sortColors75S2(nums []int) {
	step := 0
	index := 0
	for step != len(nums) {
		if nums[index] == 2 {
			// nums[:index] = nums[:index+1]
			// nums[len(nums)-1] = 2
			nums = append(append(nums[:index], nums[index+1:]...), 2)
		} else if nums[index] == 0 {
			// nums[1:index] = nums[0:index]
			// nums[0] = 0
			nums = append([]int{0}, append(nums[:index], nums[index+1:]...)...)
			index++
		} else {
			index++
		}
		step++
	}
	fmt.Println(nums)
}
