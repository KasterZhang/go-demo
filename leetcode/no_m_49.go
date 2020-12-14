package leetcode

import (
	"strconv"
	"strings"
)

//DONE
func groupAnagramsS1(strs []string) [][]string {
	dict := make(map[string][]string)
	result := make([][]string, 0, len(strs))
	for _, word := range strs {
		index := sp(word)
		if arr, ok := dict[index]; ok {
			dict[index] = append(arr, word)
		} else {
			dict[index] = []string{word}
		}
	}

	for _, arr := range dict {
		result = append(result, arr)
	}
	return result
}

func sp(word string) string {
	letter := make([]int, 26)
	for _, r := range word {
		letter[r-'a']++
	}
	var sb strings.Builder
	for _, c := range letter {
		sb.WriteString(strconv.Itoa(c))
		sb.WriteRune(' ')
	}
	return sb.String()
}

//DONE
func groupAnagramsS2(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		letter := [26]int{}
		for _, r := range str {
			letter[r-'a']++
		}
		if arr, ok := m[letter]; ok {
			m[letter] = append(arr, str)
		} else {
			m[letter] = []string{str}
		}
	}
	result := make([][]string, 0, len(strs))
	for _, arr := range m {
		result = append(result, arr)
	}
	return result
}
