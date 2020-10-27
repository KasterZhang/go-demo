package leetcode

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE 中序
func getMinimumDifferenceS1(root *TreeNode) int {
	result := math.MaxInt64
	var lrn func(node *TreeNode)
	var pre *TreeNode
	lrn = func(node *TreeNode) {
		if node == nil || result == 0 {
			return
		}
		lrn(node.Left)
		if pre != nil {
			temp := node.Val - pre.Val
			if result > temp {
				result = temp
			}
		}
		pre = node
		lrn(node.Right)
	}
	lrn(root)
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE 先序
func getMinimumDifferenceS2(root *TreeNode) int {
	stack := []*TreeNode{root}
	result := math.MaxInt64
	for len(stack) != 0 {
		pop := stack[0]
		stack = stack[1:]
		leftMin, rightMin := math.MaxInt64, math.MaxInt64
		if pop.Left != nil {
			node := pop.Left
			for node != nil {
				leftMin = pop.Val - node.Val
				if leftMin == 0 {
					return 0
				}
				node = node.Right
			}
			stack = append(stack, pop.Left)
		}

		if pop.Right != nil {
			node := pop.Right
			for node != nil {
				rightMin = node.Val - pop.Val
				if rightMin == 0 {
					return 0
				}
				node = node.Left
			}
			stack = append(stack, pop.Right)
		}
		// fmt.Println("pop", pop.Val, "leftMin", leftMin, "rightMin", rightMin)
		if leftMin <= rightMin && leftMin < result {
			result = leftMin
		}
		if rightMin < leftMin && rightMin < result {
			result = rightMin
		}
	}
	return result
}
