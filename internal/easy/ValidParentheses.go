package easy

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	if l%2 != 0 {
		return false
	}

	lib := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	parentheses := []rune{}

	for _, c := range s {
		_, open := lib[c]
		if open {
			parentheses = append(parentheses, c)
		} else if len(parentheses) != 0 {
			prev := parentheses[len(parentheses)-1]
			v := lib[prev]
			if v == c {
				parentheses = parentheses[:len(parentheses)-1]
			} else {
				return false
			}

		} else {
			return false
		}
	}

	return len(parentheses) == 0
}
