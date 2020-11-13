package leetcode

import "fmt"

//TODO
func minimumOperationsLCP19(leaves string) int {
	result := 0
	if leaves[0] != 'r' {
		result++
		leaves = leaves[1:]
	}
	if leaves[len(leaves)-1] != 'r' {
		result++
		leaves = leaves[:len(leaves)-1]
	}

	left, right, temp := cut(leaves, 'r')
	fmt.Println("temp", temp, "left", left, "right", right)
	if temp == leaves {
		if left == 0 {
			return 2
		}
		return 1
	}
	if left == 0 {
		result++
	}
	if right == 0 {
		result++
	}
	leaves = temp
	fmt.Println("result", result, "left", left, "right", right, "leaves", leaves)
	return result
}

func cut(leaves string, target byte) (int, int, string) {
	left, right, start, end, lastIndex := 0, 0, -1, -1, len(leaves)-1
	for i := range leaves {
		if start != -1 && end != -1 {
			leaves = leaves[start : end+1]
			break
		}
		if start == -1 && leaves[i] != target {
			start = i
		} else {
			left++
		}
		if end == -1 && leaves[lastIndex-i] != target {
			end = lastIndex - i
		} else {
			right++
		}
	}
	fmt.Println(start, end)
	return left, right, leaves
}
