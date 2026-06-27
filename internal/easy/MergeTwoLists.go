package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}
	var res *ListNode
	if list1.Val <= list2.Val {
		res = list1
		list1 = list1.Next
		res.Next = nil
	} else {
		res = list2
		list2 = list2.Next
		res.Next = nil
	}

	cur := res

	for {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				cur.Next = list1
				list1 = list1.Next
				cur = cur.Next
				cur.Next = nil
			} else {
				cur.Next = list2
				list2 = list2.Next
				cur = cur.Next
				cur.Next = nil
			}
		} else if list1 != nil && list2 == nil {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
			cur.Next = nil
		} else if list2 != nil && list1 == nil {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
			cur.Next = nil
		} else {
			return res
		}
	}
}
