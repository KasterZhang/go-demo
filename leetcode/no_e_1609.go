package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE
func isEvenOddTree1609(root *TreeNode) bool {
	queue := []*TreeNode{root}
	even := true
	for len(queue) != 0 {
		temp := make([]*TreeNode, 0, len(queue)*2)
		pre := 0
		for _, node := range queue {
			if node != nil {
				if even {
					fmt.Println("even", pre, node.Val)
					if node.Val%2 == 0 || node.Val <= pre {
						return false
					}
					pre = node.Val
					temp = append(temp, node.Left, node.Right)
				} else {
					fmt.Println("odd", pre, node.Val)
					if node.Val%2 == 1 || (pre != 0 && node.Val >= pre) {
						return false
					}
					pre = node.Val
					temp = append(temp, node.Left, node.Right)
				}
			}
		}
		queue = temp
		even = !even
	}
	return true
}
