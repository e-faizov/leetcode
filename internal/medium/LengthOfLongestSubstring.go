package medium

func lengthOfLongestSubstring(s string) int {
	counter := map[byte]int{}
	var maxLen, j int

	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		if counter[s[i]-'a'] > 1 {
		left:
			for j < i {
				counter[s[j]-'a']--
				j++
				if counter[s[i]-'a'] <= 1 {
					break left
				}
			}
		}
		maxLen = max(maxLen, i-j+1)
	}

	return maxLen
}
