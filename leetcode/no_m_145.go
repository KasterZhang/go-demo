package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE 基础
func primaryPostorderTraversal(root *TreeNode) []int {
	//----------------基础----------------
	var rln func(*TreeNode) []int
	rln = func(root *TreeNode) []int {
		result := []int{}
		if root == nil {
			return result
		}
		result = append(result, rln(root.Right)...)
		result = append(result, rln(root.Left)...)
		result = append(result, root.Val)
		return result
	}
	return rln(root)
}

//DONE 进阶
func juniorPostorderTraversal(root *TreeNode) []int {
	//----------------进阶----------------
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	desc := []int{}
	for {
		if len(stack) == 0 {
			break
		}
		tail := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tail != nil {
			desc = append(desc, 0)
			copy(desc[1:], desc[0:])
			desc[0] = tail.Val
			stack = append(stack, tail.Left)
			stack = append(stack, tail.Right)
		}
	}
	return desc
}
