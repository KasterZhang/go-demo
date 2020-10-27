package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//NOT GOOD
func convertBST(root *TreeNode) *TreeNode {
	if root != nil {
		rnl(root, 0)
	}
	return root
}

func rnl(root *TreeNode, tmp int) int {
	if root == nil {
		return 0
	}
	if root.Right != nil {
		tmp = rnl(root.Right, tmp)
	}
	root.Val += tmp
	if root.Left != nil {
		return rnl(root.Left, root.Val)
	}
	return root.Val
}

//BETTER
func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	var rnl2 func(*TreeNode) *TreeNode
	rnl2 = func(node *TreeNode) *TreeNode {
		if node != nil {
			rnl2(node.Right)
			node.Val += sum
			sum = node.Val
			rnl2(node.Left)
		}
		return node
	}
	return rnl2(root)
}
