package leetcode

import "fmt"

//TODO
func mostCompetitive1673(nums []int, k int) []int {
	l, result := len(nums), make([]int, k)
	index := 0
	for i := 0; i < k; i++ {
		min, count, tempIndex := -1, 0, index
		fmt.Println("start from", index, "len", l, "i", i)
		for j, n := range nums[index:] {
			if l-j-index < k-i-count {
				//剩余未遍历元素个数为l-index-j-1;本次会插入元素个数为count个 剩余需要插入的个数为k-i-1-count个
				fmt.Println("this", j+index, "left", l-index-j, "target", k-i, "test", len(nums[tempIndex+1:]))
				break
			}
			if (min == -1 || n < min) && l-index-j >= k-i {
				fmt.Println("find min", n, "at", j+index)
				count = 1
				min = n
				tempIndex = j + index
			} else if n == min {
				count++
				tempIndex = j + index
				fmt.Println("count", count, "at", j+index)
			}
		}
		index = tempIndex + 1
		fmt.Println("end at", index)
		fmt.Println("min", min, "count", count)
		fmt.Println("-------------")
		result[i] = min
		for count > 1 {
			i++
			result[i] = min
			count--
		}
	}
	return result
}
