package easy

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head.Next
	fast := head.Next
	for {
		if slow.Next == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
}
