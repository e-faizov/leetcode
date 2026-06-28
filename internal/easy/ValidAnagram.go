package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := len(s)
	arr := [28]int{}
	for _, c := range s {
		arr[c-'a']++
	}
	for _, c := range t {
		pos := c - 'a'
		if arr[pos] > 0 {
			arr[pos]--
			count--
		} else {
			return false
		}
	}

	return count == 0
}
