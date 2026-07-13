package medium

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	product := 1
	for i, v := range nums {
		res[i] = product
		product *= v
	}
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= product
		product *= nums[i]
	}
	return res
}
