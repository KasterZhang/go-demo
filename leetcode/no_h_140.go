package leetcode

import (
	"fmt"
)

// 分隔时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。

// 输入:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// 输出:
// [
//   "cats and dog",
//   "cat sand dog"
// ]

//DONE
func wordBreak140S1(s string, wordDict []string) []string {
	result := []string{}
	for _, word := range wordDict {
		if s == word {
			result = append(result, word)
		}
		if len(s) >= len(word) && s[:len(word)] == word {
			fmt.Println(s, word)
			temp := wordBreak140S1(s[len(word):], wordDict)
			for _, v := range temp {
				if len(v) != 0 {
					result = append(result, word+" "+v)
				}
			}
		}
	}
	return result
}

//TODO 剪枝优化
func wordBreak140S2(s string, wordDict []string) []string {
	result := []string{}

	return result
}
