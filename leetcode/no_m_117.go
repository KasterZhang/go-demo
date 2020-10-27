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
func primaryConnect(root *Node) *Node {
	if root == nil {
		return root
	}
	//----------------基础----------------
	temp := []*Node{root}
	for {
		if len(temp) == 0 {
			break
		}
		nextTemp := []*Node{}
		for i, n := range temp {
			if i < len(temp)-1 {
				n.Next = temp[i+1]
			}
			if n.Left != nil {
				nextTemp = append(nextTemp, n.Left)
			}
			if n.Right != nil {
				nextTemp = append(nextTemp, n.Right)
			}
		}
		temp = nextTemp
	}
	return root
}

//DONE 进阶
func juniorConnect(root *Node) *Node {
	if root == nil {
		return root
	}
	//----------------进阶----------------
	header := new(Node)
	header.Next = root
out:
	for {

		nextLineHeader := new(Node)
		nextLine := nextLineHeader

	inner:
		for {

			if header.Left != nil {
				nextLineHeader.Next = header.Left
				nextLineHeader = nextLineHeader.Next
			}
			if header.Right != nil {
				nextLineHeader.Next = header.Right
				nextLineHeader = nextLineHeader.Next
			}

			header = header.Next
			if header == nil {
				break inner
			}

		}
		header = nextLine
		if nextLine.Next == nil {
			break out
		}
	}
	return root
}
