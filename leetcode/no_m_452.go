package leetcode

import (
	"sort"
)

//DONE
func findMinArrowShots452(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Sort(Points452(points))

	result := 1
	pre := points[0]
	for _, point := range points[1:] {
		if merged, ok := merge452(pre, point); ok {
			pre = merged
		} else {
			result++
			pre = point
		}
	}
	return result
}

//Points452 452
type Points452 [][]int

func (p Points452) Len() int {
	return len(p)
}

func (p Points452) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p Points452) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func merge452(pa, pb []int) ([]int, bool) {
	if pa[0] <= pb[0] && pb[1] <= pa[1] {
		return pb, true
	} else if pb[0] <= pa[0] && pa[1] <= pb[1] {
		return pa, true
	} else if pa[0] <= pb[0] && pa[1] >= pb[0] {
		return []int{pb[0], pa[1]}, true
	} else if pb[0] <= pa[0] && pb[1] >= pa[0] {
		return []int{pa[0], pb[1]}, true
	} else {
		return []int{}, false
	}
}
