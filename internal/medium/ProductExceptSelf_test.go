package medium

import (
	"slices"
	"testing"
)

type ProductExceptSelfData struct {
	data []int
	want []int
}

func TestProductExceptSelf(t *testing.T) {
	datas := []ProductExceptSelfData{
		{
			data: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			data: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, testData := range datas {
		result := productExceptSelf(testData.data)
		if !slices.Equal(result, testData.want) {
			t.Errorf("test fail want: %v, result: %v", testData.want, result)
		}
	}
}
