package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//DONE
func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	desc, slow, fast := head, head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		temp := *slow
		temp.Next = desc
		desc = &temp
	}
	if fast.Next == nil {
		desc = desc.Next
	}
	slow = slow.Next
	if desc.Val != slow.Val {
		return false
	}
	for slow != nil {
		fmt.Println(slow.Val, desc.Val)
		if desc.Val != slow.Val {
			return false
		}
		slow = slow.Next
		desc = desc.Next
	}
	return true
}
