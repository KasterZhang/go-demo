package leetcode

//TreeNode leetcode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Node leetcode struct
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//ListNode leetcode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func get(node *TreeNode) []interface{} {
	ch := make(chan *TreeNode, 10000)
	ch <- node
	result := []interface{}{}
	for n := range ch {
		if n != nil {
			result = append(result, n.Val)
			ch <- n.Left
			ch <- n.Right
		} else {
			result = append(result, "null")
		}
	}
	return result
}
