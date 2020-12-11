package leetcode

import (
	"strconv"
	"strings"
)

//DONE
func splitIntoFibonacci842(S string) []int {
	index0, index1, l := 1, 2, len(S)
	if l < 3 {
		return []int{}
	}

	for index0 <= l/2 {
		result := make([]int, 2, len(S))
		result[0], _ = strconv.Atoi(S[:index0])
		result[1], _ = strconv.Atoi(S[index0:index1])
		temp := S[index1:]
		target := result[0] + result[1]
		for len(temp) != 0 {
			ll := len(result)
			targetStr := strconv.Itoa(target)
			if strings.HasPrefix(temp, targetStr) {
				if target > 1<<31-1 {
					break
				}
				result = append(result, target)
				target = 2*target - result[ll-2]
				temp = temp[len(targetStr):]
			} else {
				break
			}
		}
		if len(temp) == 0 {
			return result
		}

		if l > 2*index1-index0 && S[index0] != '0' {
			index1++
		} else {
			if S[0] == '0' {
				break
			}
			index0++
			index1 = index0 + 1
		}
	}
	return []int{}
}
