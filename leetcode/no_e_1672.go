package leetcode

//DONE
func maximumWealth1672(accounts [][]int) int {
	max := 0
	for i := range accounts {
		temp := 0
		for _, n := range accounts[i] {
			temp += n
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
