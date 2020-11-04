package leetcode

//DONE
func minHeightShelves1105S1(books [][]int, shelf_width int) int {
	dp := []int{0}
	for i, book := range books {
		minH, tempH, width := book[1]+dp[i], book[1], shelf_width-book[0]
		for j := len(dp) - 1; j > 0; j-- {
			if books[j][0] <= width {
				width -= books[j][0]
			} else {
				break
			}
			if books[j][1] > tempH {
				tempH = books[j][1]
			}
			if minH > dp[j]+tempH {
				minH = dp[j] + tempH
			}
		}
		dp = append(dp, minH)
	}
	return dp[len(books)]
}
