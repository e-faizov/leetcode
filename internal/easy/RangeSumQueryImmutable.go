package easy

type NumArray struct {
	sNums []int
}

func Constructor(nums []int) NumArray {
	sNums := make([]int, len(nums))
	var sum int
	for i, v := range nums {
		sum += v
		sNums[i] = sum
	}
	return NumArray{
		sNums: sNums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sNums[right]
	}
	return this.sNums[right] - this.sNums[left-1]
}
