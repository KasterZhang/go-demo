package leetcode

import (
	"fmt"
	"sort"
)

//Tasks Tasks
type Tasks [][]int

func (ts Tasks) Len() int { return len(ts) }
func (ts Tasks) Less(i, j int) bool {
	if (ts[i][1]-ts[i][0] < ts[j][1]-ts[j][0]) || (ts[i][1]-ts[i][0] == ts[j][1]-ts[j][0] && ts[i][1] < ts[j][1]) {
		return true
	}
	return false
}
func (ts Tasks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

//DONE
func minimumEffort(tasks [][]int) int {
	sort.Sort(Tasks(tasks))
	fmt.Println(tasks)
	result := 0
	for i, t := range tasks {
		if i == 0 {
			result += t[1]
		} else if result+t[0] >= t[1] {
			result += t[0]
		} else {
			result += t[1]
		}
	}
	return 0
}
