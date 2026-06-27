package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	{
		data := []string{"flower", "flow", "flight"}
		want := "fl"
		result := longestCommonPrefix(data)
		if result != want {
			t.Errorf("test fail. '%v' != '%v'", want, result)
		}
	}
	{
		data := []string{"dog", "racecar", "car"}
		want := ""
		result := longestCommonPrefix(data)
		if result != want {
			t.Errorf("test fail. '%v' != '%v'", want, result)
		}
	}
	{
		data := []string{"ab", "a"}
		want := "a"
		result := longestCommonPrefix(data)
		if result != want {
			t.Errorf("test fail. '%v' != '%v'", want, result)
		}
	}
}
