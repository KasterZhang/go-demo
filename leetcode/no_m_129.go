package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE
func sumNumbers129S1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	result := 0
	for len(queue) != 0 {
		pop := queue[0]
		queue = queue[1:]
		lastNode := pop[len(pop)-1]
		if lastNode.Left == nil && lastNode.Right == nil {
			temp := 0
			last := len(pop) - 1
			for i, v := range pop {
				temp += v.Val * int(math.Pow10(last-i))
			}
			result += temp
		}
		if lastNode.Left != nil {
			temp := make([]*TreeNode, len(pop)+1)
			copy(temp, pop)
			temp[len(temp)-1] = lastNode.Left
			queue = append(queue, temp)
		}
		if lastNode.Right != nil {
			temp := make([]*TreeNode, len(pop)+1)
			copy(temp, pop)
			temp[len(temp)-1] = lastNode.Right
			queue = append(queue, temp)
		}
	}
	return result
}

//DONE
func sumNumbers129S2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	resultArr := []int{root.Val}
	result := 0
	for len(queue) != 0 {
		pop := queue[0]
		popVal := resultArr[0]
		queue = queue[1:]
		resultArr = resultArr[1:]
		if pop.Left == nil && pop.Right == nil {
			result += popVal
		}
		if pop.Left != nil {
			queue = append(queue, pop.Left)
			resultArr = append(resultArr, popVal*10+pop.Left.Val)
		}
		if pop.Right != nil {
			queue = append(queue, pop.Right)
			resultArr = append(resultArr, popVal*10+pop.Right.Val)
		}
	}
	return result
}
