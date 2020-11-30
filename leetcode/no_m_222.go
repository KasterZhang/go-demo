package leetcode

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE
func countNodes222S1(root *TreeNode) int {

	result := 0

	queue := make([]*TreeNode, 0, 1)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		result += len(queue)
		next := make([]*TreeNode, 0, len(queue))
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		queue = next
	}
	return result
}

//DONE
func countNodes222S2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		node := root
		for h := level - 1; h >= 0; h-- {
			if (k>>h)&1 == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		return node == nil
	}) - 1
}
