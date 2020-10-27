package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//DONE 递归
func preorderTraversalS1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	result = append(result, preorderTraversalS1(root.Left)...)
	return append(result, preorderTraversalS1(root.Right)...)
}

//DONE 迭代
func preorderTraversalS2(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop == nil {
			continue
		}
		result = append(result, pop.Val)
		stack = append(stack, pop.Right)
		stack = append(stack, pop.Left)
	}
	return result
}
