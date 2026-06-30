package medium

func groupAnagrams(strs []string) [][]string {
	tmp := map[[26]byte][]string{}
	for _, str := range strs {
		lib := [26]byte{}
		for _, c := range str {
			lib[c-'a']++
		}
		tmp[lib] = append(tmp[lib], str)
	}
	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}
