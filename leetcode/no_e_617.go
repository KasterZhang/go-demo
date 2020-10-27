package leetcode

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return dfs(t1, t2)
}

func dfs(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = dfs(t1.Left, t2.Left)
	t1.Right = dfs(t1.Right, t2.Right)
	return t1
}
