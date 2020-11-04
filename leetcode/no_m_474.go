package leetcode

// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
//TODO
func findMaxForm474(strs []string, m int, n int) int {
	// result := []int{}
	// for i, str := range strs {

	// }
	return 0
}

func analyse474(s string) (int, int) {
	m, n := 0, 0
	for _, c := range s {
		if c == rune(0) {
			m++
		} else {
			n++
		}
	}
	return m, n
}
