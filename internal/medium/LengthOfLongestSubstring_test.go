package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	{
		data := " "
		want := 1
		result := lengthOfLongestSubstring(data)
		if result != want {
			t.Errorf("test fail. Data: '%v' '%v' != '%v'", data, want, result)
		}
	}
	{
		data := "abcabcbb"
		want := 3
		result := lengthOfLongestSubstring(data)
		if result != want {
			t.Errorf("test fail. Data: '%v' '%v' != '%v'", data, want, result)
		}
	}
	{
		data := "bbbbb"
		want := 1
		result := lengthOfLongestSubstring(data)
		if result != want {
			t.Errorf("test fail. Data: '%v' '%v' != '%v'", data, want, result)
		}
	}
	{
		data := "pwwkew"
		want := 3
		result := lengthOfLongestSubstring(data)
		if result != want {
			t.Errorf("test fail. Data: '%v' '%v' != '%v'", data, want, result)
		}
	}
}
