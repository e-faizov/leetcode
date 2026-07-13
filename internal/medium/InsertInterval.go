package medium

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	inserted := false
	process := func(next []int) {
		if len(res) == 0 {
			res = append(res, next)
		} else {
			if res[len(res)-1][1] >= next[0] {
				res[len(res)-1][1] = max(res[len(res)-1][1], next[1])
			} else {
				res = append(res, next)
			}
		}
	}

	for _, v := range intervals {
		if !inserted && newInterval[0] < v[0] {
			inserted = true
			process(newInterval)
		}
		process(v)
	}
	if !inserted {
		process(newInterval)
	}
	return res
}
