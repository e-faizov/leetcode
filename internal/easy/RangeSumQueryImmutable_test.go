package easy

import "testing"

func TestSumRange(t *testing.T) {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	obj.SumRange(2, 5)
}
