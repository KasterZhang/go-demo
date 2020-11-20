package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//TODO 回顾
//DONE
func insertionSortList147(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	start := &ListNode{Next: head}
	index, loop := head, head.Next
	for loop != nil {
		if index.Val <= loop.Val {
			//满足设定无需移动
			index = index.Next
		} else {
			temp := start
			for temp.Next.Val <= loop.Val {
				//下一个比当前的小则继续
				temp = temp.Next
			}
			//temp的next比当前的大
			//连接sorted和剩余未循环部分
			index.Next = loop.Next
			//当前点位插入到temp之后
			loop.Next = temp.Next
			temp.Next = loop
		}
		//下一个点位指向sorted的next
		loop = index.Next
	}
	return start.Next
}
