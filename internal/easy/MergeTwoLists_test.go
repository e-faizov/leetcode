package easy

import "testing"

func makeList(data []int) *ListNode {
	var res, cur *ListNode
	for _, v := range data {
		if res == nil {
			res = &ListNode{
				Val: v,
			}
			cur = res
		} else {
			next := &ListNode{
				Val: v,
			}
			cur.Next = next
			cur = next
		}
	}
	return res
}

func checkList(data *ListNode, checkValue []int) bool {
	if data == nil && len(checkValue) == 0 {
		return true
	}
	i := 0
	for {
		if data.Val != checkValue[i] {
			return false
		}
		if data.Next != nil {
			i++
			if i >= len(checkValue) {
				return false
			}
			data = data.Next
		} else {
			if i+1 != len(checkValue) {
				return false
			} else {
				return true
			}
		}
	}
}

func TestMergeTwoLists(t *testing.T) {
	{
		data1 := makeList([]int{1, 2, 4})
		data2 := makeList([]int{1, 3, 4})

		result := mergeTwoLists(data1, data2)
		if !checkList(result, []int{1, 1, 2, 3, 4, 4}) {
			t.Error("Error")
		}
	}
	{
		data1 := makeList([]int{})
		data2 := makeList([]int{})

		result := mergeTwoLists(data1, data2)
		if !checkList(result, []int{}) {
			t.Error("Error")
		}
	}
	{
		data1 := makeList([]int{})
		data2 := makeList([]int{0})

		result := mergeTwoLists(data1, data2)
		if !checkList(result, []int{0}) {
			t.Error("Error")
		}
	}
}
