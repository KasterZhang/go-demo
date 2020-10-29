package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//DONE
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2, pre := 0, 0, 0
	count1, count2 := l1, l2
	for count1 != nil {
		len1++
		count1 = count1.Next
	}

	for count2 != nil {
		len2++
		count2 = count2.Next
	}
	var result *ListNode
	kk := result
	over10 := false
	for len1 != 0 || len2 != 0 || over10 {
		temp := 0
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
			len1--
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
			len2--
		}
		if over10 {
			temp++
			over10 = false
		}

		if temp >= 10 {
			over10 = true
			pre = temp - 10
		} else {
			pre = temp
		}
		node := new(ListNode)
		node.Val = pre
		if kk == nil {
			result = node
			kk = result
		} else {
			kk.Next = node
			kk = kk.Next
		}
	}
	return result
}
