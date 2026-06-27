package easy

import "testing"

func TestRomanToInt(t *testing.T) {
	{
		data := "III"
		want := 3
		result := romanToInt(data)
		if result != want {
			t.Errorf("test fail. data '%v' '%v' != '%v'", data, want, result)
		}
	}
	{
		data := "LVIII"
		want := 58
		result := romanToInt(data)
		if result != want {
			t.Errorf("test fail. data '%v' '%v' != '%v'", data, want, result)
		}
	}
	{
		data := "MCMXCIV"
		want := 1994
		result := romanToInt(data)
		if result != want {
			t.Errorf("test fail. data '%v' '%v' != '%v'", data, want, result)
		}
	}
}
