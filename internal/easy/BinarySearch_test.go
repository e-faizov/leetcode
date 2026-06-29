package easy

import "testing"

func TestSearch(t *testing.T) {
	{
		data := []int{-1, 0, 3, 5, 9, 12}
		target := 3
		want := 2
		result := search(data, target)
		if result != want {
			t.Errorf("test fail. Target: '%v' '%v' != '%v'", target, want, result)
		}
	}
	{
		data := []int{-1, 0, 3, 5, 9, 12}
		target := 2
		want := -1
		result := search(data, target)
		if result != want {
			t.Errorf("test fail. Target: '%v' '%v' != '%v'", target, want, result)
		}
	}
}
