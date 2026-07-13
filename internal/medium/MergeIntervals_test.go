package medium

import (
	"slices"
	"testing"
)

type MergeIntervalsData struct {
	data [][]int
	want [][]int
}

func TestMergeIntervals(t *testing.T) {
	testDatas := []MergeIntervalsData{
		{
			data: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			data: [][]int{
				{4, 7},
				{1, 4},
			},
			want: [][]int{
				{1, 7},
			},
		},
		{
			data: [][]int{
				{1, 4},
				{4, 5},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			data: [][]int{
				{1, 4},
				{2, 3},
			},
			want: [][]int{
				{1, 4},
			},
		},
	}
	for _, v := range testDatas {
		result := mergeIntervals(v.data)

		if !slices.EqualFunc(result, v.want, func(e1, e2 []int) bool {
			return slices.Equal(e1, e2)
		}) {
			t.Errorf("test fail want: %v, result: %v", v.want, result)
		}

	}

}
