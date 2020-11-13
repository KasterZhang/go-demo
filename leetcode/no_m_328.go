package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//TODO 指针用法优化
//DONE
func oddEvenList328(head *ListNode) *ListNode {
	loop := head
	i := 0
	tempOdd, tempEven := new(ListNode), new(ListNode)
	odd, even := tempOdd, tempEven
	for loop != nil {
		i++
		temp := *loop
		if i&1 == 0 {
			tempEven.Next = &temp
			tempEven = tempEven.Next
			if tempEven != nil {
				tempEven.Next = nil
			}
		} else {
			tempOdd.Next = &temp
			tempOdd = tempOdd.Next
			if tempOdd != nil {
				tempOdd.Next = nil
			}
		}
		loop = loop.Next
	}
	tempOdd.Next = even.Next
	return odd.Next
}
