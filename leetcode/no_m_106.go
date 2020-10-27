package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	if len(inorder) == 1 && len(postorder) == 1 {
		return root
	}
	rootLocation := find(inorder, root.Val)
	left := inorder[:rootLocation]
	root.Left = buildTree(left, postorder[:len(left)])
	right := inorder[rootLocation+1:]
	root.Right = buildTree(right, postorder[len(postorder)-len(right)-1:len(postorder)-1])
	return root
}

func find(arr []int, target int) int {
	for i, number := range arr {
		if number == target {
			return i
		}
	}
	return -1
}
