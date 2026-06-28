package easy

import "testing"

func TestXxx(t *testing.T) {
	{
		data1 := "anagram"
		data2 := "nagaram"
		want := true
		result := isAnagram(data1, data2)
		if result != want {
			t.Errorf("test fail. '%v' != '%v'", want, result)
		}
	}
	{
		data1 := "rat"
		data2 := "car"
		want := false
		result := isAnagram(data1, data2)
		if result != want {
			t.Errorf("test fail. '%v' != '%v'", want, result)
		}
	}

}
