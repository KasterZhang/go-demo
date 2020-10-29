package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//TODO
func removeNthFromEnd19(head *ListNode, n int) *ListNode {
	result := head
	var pre, tail *ListNode = result, result
	tailCount := n - 1
	preCount := n + 1
	for head != nil {
		if head == nil {
			break
		}
		if tailCount == 0 {
			tail = tail.Next
		} else {
			tailCount--
		}
		if preCount == 0 {
			pre = pre.Next
		} else {
			preCount--
		}
		head = head.Next
	}
	if preCount != 0 {
		return result.Next
	}
	pre.Next = tail

	return result
}
