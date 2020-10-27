package leetcode

//DONE
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	step := 0
	catch := false
	for fast.Next != nil && fast.Next.Next != nil {
		if slow == fast {
			catch = true
			break
		}
		step++
	}
	if !catch {
		return nil
	}
	result := head
a:
	for {
		for i := 0; i < step*2; i++ {
			if slow == result {
				break a
			}
			slow = slow.Next

		}
		result = result.Next
	}
	return result
}
