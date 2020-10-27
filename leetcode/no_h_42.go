package leetcode

//执行用时：
// 232 ms, 在所有 Go 提交中击败了5.46%的用户
// 内存消耗：
// 3.2 MB, 在所有 Go 提交中击败了5.07%的用户
//DONE BAD TODO REVIEW
func trap42(height []int) int {
	max := 0
	for _, h := range height {
		if h > max {
			max = h
		}
	}
	hkey := make([]int, max+1)
	result := 0
	for location, h := range height {
		// fmt.Printf("------%d-------\n", h)
		// fmt.Println(hkey[:h+1])
		for tempH, preLocation := range hkey[:h+1] {
			if tempH != 0 && preLocation != 0 {
				preLocation = preLocation - 1
				// fmt.Println(location, h, preLocation, tempH)
				// fmt.Println("result add", location-preLocation-1)
				result += location - preLocation - 1
			}
			hkey[tempH] = location + 1
		}
	}
	return result
}
