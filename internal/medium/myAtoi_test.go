package medium

import "testing"

func TestMyAtoi(t *testing.T) {
	{
		data := "+-12"
		want := 0
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "+1"
		want := 1
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "-91283472332"
		want := MinInt32
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "42"
		want := 42
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := " -042"
		want := -42
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "1337c0d3"
		want := 1337
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "0-1"
		want := 0
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "words and 987"
		want := 0
		result := myAtoi(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}

}
