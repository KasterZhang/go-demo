package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
//DONE
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lnr func(*TreeNode, *TreeNode, *TreeNode) *TreeNode
	lnr = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		findL := lnr(root.Left, p, q)
		if root.Val == p.Val {
			return p
		}
		if root.Val == q.Val {
			return q
		}
		findR := lnr(root.Right, p, q)
		if findL != nil && findR != nil && findL != findR {
			return root
		}
		if findL != nil {
			return findL
		}
		return findR
	}

	return lnr(root, p, q)
}
