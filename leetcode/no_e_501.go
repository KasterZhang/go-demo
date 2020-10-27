package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	return []int{}
}

//TODO 基础
func primaryFindMode(root *TreeNode) []int {
	//----------------基础----------------
	if root == nil {
		return []int{}
	}
	pre := primaryFindMode(root.Left)
	if len(pre) > 0 && root.Val == pre[len(pre)-1] {
		pre = append(pre, root.Val)
	}
	primaryFindMode(root.Right)

	return nil
}

//TODO 进阶
func juniorFindMode(root *TreeNode) []int {
	//----------------进阶----------------
	return nil
}
