package medium

import (
	"slices"
	"testing"
)

type InsertIntervalData struct {
	data   [][]int
	insert []int
	want   [][]int
}

func TestIntervalInterval(t *testing.T) {
	testDatas := []InsertIntervalData{
		{
			data: [][]int{
				{1, 3},
				{6, 9},
			},
			insert: []int{2, 5},
			want: [][]int{
				{1, 5},
				{6, 9},
			},
		},
		//Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
		//Output: [[1,2],[3,10],[12,16]]
		{
			data: [][]int{
				{1, 2},
				{3, 5},
				{6, 7},
				{8, 10},
				{12, 16},
			},
			insert: []int{4, 8},
			want: [][]int{
				{1, 2},
				{3, 10},
				{12, 16},
			},
		},
		{
			data:   [][]int{},
			insert: []int{5, 7},
			want: [][]int{
				{5, 7},
			},
		},
		{
			data: [][]int{
				{1, 5},
			},
			insert: []int{2, 7},
			want: [][]int{
				{1, 7},
			},
		},
	}
	for _, v := range testDatas {
		result := insertInterval(v.data, v.insert)

		if !slices.EqualFunc(result, v.want, func(e1, e2 []int) bool {
			return slices.Equal(e1, e2)
		}) {
			t.Errorf("test fail want: %v, result: %v", v.want, result)
		}

	}

}
