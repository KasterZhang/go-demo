package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
//DONE
func primary116(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for {
		if len(queue) == 0 {
			break
		}
		temp := []*Node{}
		for i, node := range queue {
			if node.Left != nil {
				temp = append(temp, node.Left, node.Right)
			}
			if i != len(queue)-1 {
				node.Next = queue[i+1]
			}
		}
		queue = temp
	}
	return root
}

func junior116(root *Node) *Node {
	header := new(Node)
	header.Next = root
	temp116(header)
	return root
}

func temp116(header *Node) {
	if header.Next == nil {
		return
	}
	nextHeader := new(Node)
	nextTemp := new(Node)
	for {
		if header.Next == nil {
			break
		}
		if header.Next.Left != nil {
			if nextHeader.Next == nil {
				nextHeader.Next = header.Next.Left
			}
			nextTemp.Next = header.Next.Left
			header.Next.Left.Next = header.Next.Right
			nextTemp = header.Next.Right
		}
		header = header.Next
	}
	temp116(nextHeader)
}
