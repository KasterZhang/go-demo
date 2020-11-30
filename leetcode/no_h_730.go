package leetcode

import (
	"fmt"
)

//给定一个字符串 S，找出 S 中不同的非空回文子序列个数，并返回该数字与 10^9 + 7 的模。
// 通过从 S 中删除 0 个或多个字符来获得子序列。
// 如果一个字符序列与它反转后的字符序列一致，那么它是回文字符序列。
// 如果对于某个  i，A_i != B_i，那么 A_1, A_2, ... 和 B_1, B_2, ... 这两个字符序列是不同的。
// 字符串 S 的长度将在[1, 1000]范围内。
// 每个字符 S[i] 将会是集合 {'a', 'b', 'c', 'd'} 中的某一个。
//TODO DP
func countPalindromicSubsequences730(S string) int {
	l := len(S)
	//初始化DP
	dp := make([][]int, 'd'-'a'+1)
	for i := range dp {
		dp[i] = make([]int, l)
	}
	last := []int{-1, -1, -1, -1}
	for i, n := range S {
		this := n - 'a'
		fmt.Println("this", string(n))
		for letter, cell := range dp {
			if rune(letter) == this {
				temp := 1
				for j := 0; j < 4; j++ {
					if j != letter && last[this] != -1 {
						//error
						temp += dp[j][i-1] - dp[j][last[this]]
					}
				}
				if i != 0 {
					cell[i] = cell[i-1]
				}
				cell[i] += temp
			} else {
				if i != 0 {
					cell[i] = cell[i-1]
				}
			}
		}
		last[this] = i
		fmt.Println("-------------")
	}

	result := 0
	for _, cell := range dp {
		result += cell[l-1]
		// fmt.Println("add",cell[l-1],"result", result)
	}
	fmt.Println(dp)
	return result % 1000000007
}
