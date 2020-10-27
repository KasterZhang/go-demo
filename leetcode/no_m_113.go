package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE
func pathSum(root *TreeNode, sum int) [][]int {
	var nrl func(*TreeNode, int, []int) [][]int
	nrl = func(root *TreeNode, sum int, path []int) [][]int {
		if root == nil {
			return [][]int{}
		}
		result := [][]int{}
		tmp := append(path, root.Val)

		if root.Val == sum && root.Left == nil && root.Right == nil {
			tmpR := make([]int, len(tmp))
			copy(tmpR, tmp)
			result = append(result, tmpR)
		} else {
			leftR := nrl(root.Left, sum-root.Val, tmp)
			rightR := nrl(root.Right, sum-root.Val, tmp)
			result = append(result, append(leftR, rightR...)...)

		}
		return result
	}
	result := nrl(root, sum, []int{})
	return result
}
