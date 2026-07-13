package medium

import (
	"slices"
)

func mergeIntervals(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(e1, e2 []int) int {
		if e1[0] < e2[0] {
			return -1
		} else if e1[0] > e2[0] {
			return 1
		} else {
			if e1[1] < e2[1] {
				return -1
			} else if e1[1] > e2[1] {
				return 1
			} else {
				return 0
			}
		}
	})
	res := make([][]int, 0, len(intervals))
	for i, v := range intervals {
		if i != 0 && res[len(res)-1][1] >= v[0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], v[1])
		} else {
			res = append(res, v)
		}
	}
	return res
}
