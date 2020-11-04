package leetcode

import (
	"fmt"
)

func sumOfDistancesInTree834(N int, edges [][]int) []int {
	for i := range edges {
		if edges[i][0] > edges[i][1] {
			edges[i] = []int{edges[i][1], edges[i][0]}
		}
	}

	var qs func(nums [][]int, startOrEnd int) [][]int
	qs = func(nums [][]int, startOrEnd int) [][]int {
		if len(nums) == 0 {
			return [][]int{}
		}
		left, right, base := [][]int{}, [][]int{}, nums[0]
		for i := range nums {
			if base[startOrEnd] > nums[i][startOrEnd] {
				left = append(left, nums[i])
			} else {
				right = append(right, nums[i])
			}
		}
		left = qs(left, startOrEnd)
		right = qs(right, startOrEnd)
		return append(append(left, base), right...)
	}
	// edgesStart := qs(edges, 0)
	edgesEnd := qs(edges, 1)

	// fmt.Println(edgesStart)
	fmt.Println(edgesEnd)

	var tf func(target, interval, startOrend int, edges [][]int) int
	tf = func(target, interval, startOrend int, edges [][]int) int {
		interval++
		count := 0
		for i, edge := range edges {
			if edge[startOrend] == target {
				count += interval
				next := edge[1-startOrend]
				count += tf(next, interval, startOrend, edges[i:])
			}
		}
		return count
	}

	result := []int{}
	for i := 0; i < N; i++ {
		interval := 0
		temp := 0
		// temp += tf(i, interval, 0, edgesStart)
		// fmt.Println("N", i, "start", tf(i, interval, 0, edgesStart))
		temp += tf(i, interval, 1, edgesEnd)
		fmt.Println("N", i, "end", tf(i, interval, 1, edgesEnd))
		result = append(result, temp)
		fmt.Println("--------------------")
	}
	return result
}
