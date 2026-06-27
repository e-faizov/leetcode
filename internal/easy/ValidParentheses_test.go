package easy

import "testing"

func TestValidParentheses(t *testing.T) {
	{
		data := ")("
		want := false
		result := isValid(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "()"
		want := true
		result := isValid(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "()[]{}"
		want := true
		result := isValid(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "(]"
		want := false
		result := isValid(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "([])"
		want := true
		result := isValid(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}
	{
		data := "([)]"
		want := false
		result := isValid(data)
		if result != want {
			t.Errorf("Error data: '%v', want: '%v', get: '%v'", data, want, result)
		}
	}

}
