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
//TODO DP2维
func countPalindromicSubsequences730S1(S string) int {
	l := len(S)
	//初始化DP
	dps := make([][]int, 'd'-'a'+1)
	for i := range dps {
		dps[i] = make([]int, l)
	}
	dps[S[l-1]-'a'][l-1] = 1
	for i := l - 2; i >= 0; i-- {
		tempLetter := S[i] - 'a'
		dps[tempLetter][i] = 1
		for letter := range dps {
			if letter == int(tempLetter) {
				flag := false
				for j := range dps[letter][i+1:] {
					if S[i+1+j]-'a' == tempLetter {
						flag = true
						dps[letter][i+1+j] = 2 + dps[0][i+j] + dps[1][i+j] + dps[2][i+j] + dps[3][i+j]
					} else {
						if flag {
							dps[letter][i+1+j] = dps[letter][i+j]
						} else {
							dps[letter][i+1+j]++
						}
					}
				}
			}
			if letter == 0 {
				fmt.Println(dps[letter])
			}
		}
	}
	result := 0
	for i := range dps {
		result += dps[i][l-1]
	}
	// fmt.Println(dps)
	return result % 1000000007
}

//DONE DP3维
func countPalindromicSubsequences730S2(S string) int {
	l := len(S)
	dp := make([][][]int, 'd'-'a'+1)
	for i := range dp {
		cell := make([][]int, l)
		for j := range cell {
			cell[j] = make([]int, l)
		}
		dp[i] = cell
	}
	dp[S[l-1]-'a'][l-1][l-1] = 1

	for i := l - 2; i >= 0; i-- {
		dp[S[i]-'a'][i][i] = 1

		tempLetter := S[i] - 'a'
		for letter := range dp {
			if letter == int(tempLetter) {
				flag := false
				for index := range dp[letter][i][i+1:] {
					if tempLetter == S[i+1+index]-'a' {
						flag = true
						dp[letter][i][i+1+index] = (2 + dp[0][i+1][i+index] + dp[1][i+1][i+index] + dp[2][i+1][i+index] + dp[3][i+1][i+index]) % 1000000007
					} else {
						if flag {
							dp[letter][i][i+1+index] = dp[letter][i][i+index]
						} else {
							dp[letter][i][i+1+index] = (dp[letter][i+1][i+1+index] + 1) % 1000000007
						}
					}
				}
			} else {
				for index := range dp[letter][i][i+1:] {
					dp[letter][i][i+1+index] = dp[letter][i+1][i+1+index]
				}
			}
		}
	}

	result := 0
	for i := range dp {
		result += dp[i][0][l-1]
	}
	return result % 1000000007
}
func printDP(dps [][][]int) {
	for letter, dp := range dps {
		fmt.Println("this", string(rune(letter+'a')))
		for i := range dp {
			for j := range dp[i] {
				fmt.Printf("%d ", dp[i][j])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n-----------\n")
	}
}
