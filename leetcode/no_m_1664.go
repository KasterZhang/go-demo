package leetcode

//DONE
func waysToMakeFair(nums []int) int {
	result, odd, even := 0, 0, 0
	for i, n := range nums {
		if i&1 == 0 {
			odd += n
		} else {
			even += n
		}
	}
	newOdd, newEven := 0, 0
	for i, n := range nums {
		if i&1 == 0 {
			odd -= n
		} else {
			even -= even
		}
		if odd+newOdd == even+newEven {
			result++
		}

		if i&1 == 0 {
			newEven += n
		} else {
			newOdd += n
		}
	}
	return result
}
