package leetcode

import (
	"sort"
	"strings"
)

// Arr767 Arr767
type Arr767 [][]int

func (a Arr767) Len() int { return len(a) }
func (a Arr767) Less(i, j int) bool {
	if a[i][0] > a[j][0] {
		return true
	}
	return false
}
func (a Arr767) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//DONE
func reorganizeString767(S string) string {
	m := make(map[rune]int)
	for _, r := range S {
		if count, ok := m[r]; ok {
			m[r] = count + 1
		} else {
			m[r] = 1
		}
	}
	var arr Arr767 = make([][]int, len(m))
	for r, v := range m {
		arr[r-'a'] = []int{int(r - 'a'), v}
	}
	sort.Sort(arr)
	var sb strings.Builder
	var pre rune
	for i := 0; i < len(S); i++ {
		for j, node := range arr {
			if i == 0 || rune(node[0]) != pre {
				pre = rune(node[0])
				sb.WriteRune(rune(node[0] + 'a'))
				// this := []int{node[0], node[1] - 1}
				arr[j][1] = node[1] - 1
				if arr[j][1] != 0 {
					for k, n := range arr[j+1:] {
						if n[1] > arr[j+k][1] {
							arr[j+k], arr[j+k+1] = arr[j+k+1], arr[j+k]
						} else {
							break
						}
					}
				} else {
					arr = append(arr[:j], arr[j+1:]...)
				}
				break
			}
		}
	}

	return sb.String()
}
