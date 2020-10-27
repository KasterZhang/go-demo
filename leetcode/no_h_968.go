package leetcode

import "fmt"

//DONE
func minCameraCover(root *TreeNode) int {
	//后序遍历
	count, need, _ := lrn(root, true)
	fmt.Println("has child need camera")
	if count == 0 {
		count++
	} else if need && root.Val == 0 {
		count++
	}
	return count
}

//return count childIsLeaf needCamera
func lrn(root *TreeNode, isRoot bool) (int, bool, bool) {
	childrenIsLeaf := false
	ImLeaf := false
	needCamera := true
	count := 0
	childrenNeedCamera := false
	if root.Left != nil {
		// fmt.Println("left child find")
		tempCount, leftChildisLeaf, leftNeedCamera := lrn(root.Left, false)
		childrenIsLeaf = childrenIsLeaf || leftChildisLeaf
		childrenNeedCamera = childrenNeedCamera || leftNeedCamera
		count += tempCount
	}
	if root.Right != nil {
		// fmt.Println("right child find")
		tempCount, rightChildisLeaf, rightNeedCamera := lrn(root.Right, false)
		childrenIsLeaf = childrenIsLeaf || rightChildisLeaf
		childrenNeedCamera = childrenNeedCamera || rightNeedCamera
		count += tempCount
	}
	// fmt.Println("childrenIsLeaf:", childrenIsLeaf)
	if root.Left == nil && root.Right == nil {
		//leaf
		ImLeaf = true
		fmt.Println("leaf,parrent must has camera")
		return count, true, true
	} else if ((root.Left != nil && root.Left.Val == 1) || (root.Right != nil && root.Right.Val == 1)) && !childrenIsLeaf && !childrenNeedCamera {
		// fmt.Println("chidren has camera")
		needCamera = false
		fmt.Println("no need camera")
	} else if !childrenIsLeaf && !childrenNeedCamera && !isRoot {
		//孩子不是叶子 且孩子不需要camera 且不为根节点
		fmt.Println("孩子不是叶子 且孩子不需要camera 且不为根节点")
		return count, false, true
		//has parrent?
	} else {
		root.Val = 1
		fmt.Println("need camera")
		fmt.Println("install", root.Val)
		needCamera = false
		count++
	}
	return count, ImLeaf, needCamera
}
