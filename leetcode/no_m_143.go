package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//DONE
func reorderList143S1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	desc := new(ListNode)
	count := 0
	tempHead := *head
	for {
		count++
		tmp := tempHead
		tmp.Next = nil

		temp := desc.Next
		desc.Next = &tmp
		desc.Next.Next = temp

		if tempHead.Next == nil {
			break
		}
		tempHead = *tempHead.Next
	}
	descCount := count / 2
	// fmt.Println("----------",count,descCount)
	for {
		// fmt.Println("this head",head.Val)
		nextHead := *head.Next
		// fmt.Println("next head",nextHead.Val)
		pnextHead := &nextHead
		descNode := *desc.Next
		descNode.Next = pnextHead
		desc = desc.Next
		// fmt.Println("get desc",descNode.Val)
		head.Next = &descNode
		descCount--
		if descCount == 0 {
			// fmt.Println("judge",count%2,"head",head.Val)
			if count%2 == 1 {
				head = head.Next
				head.Next.Next = nil
			} else {
				head.Next.Next = nil
			}
			break
		} else {
			head = head.Next.Next
		}
		// fmt.Println("next head2",head.Val)
		// fmt.Println("----------",count,descCount)
	}
}
//DONE
func reorderList143S2(head *ListNode) {
	//双指针
	if head == nil {
		return
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var mid *ListNode = nil
	for slow != nil {
		nextHead := slow.Next
		slow.Next = mid
		mid = slow
		slow = nextHead
	}

	for head != nil && mid != nil {
		nextHead := head.Next
		nextMid := mid.Next
		head.Next = mid
		head.Next.Next = nextHead

		mid = nextMid
		head = nextHead
	}
}
