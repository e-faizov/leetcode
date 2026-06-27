package easy

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		l := min(len(res), len(strs[i]))

	out:
		for c := 0; c < l; c++ {
			if res[c] != strs[i][c] {
				if c == 0 {
					return ""
				}
				res = res[:c]
				break out
			}
		}

		if len(strs[i]) < len(res) {
			res = strs[i]
		}
	}

	return res
}
