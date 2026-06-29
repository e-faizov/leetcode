package easy

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	start := 0
	end := l

	for {
		p := ((end - start) / 2) + start
		if nums[p] == target {
			return p
		}
		if p == start || p == end {
			return -1
		}
		if nums[p] > target {
			end = p
		} else {
			start = p
		}
	}
}
