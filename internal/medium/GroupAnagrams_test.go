package medium

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	{
		data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		/*want := [][]string{
			{"bat"},
			{"nat", "tan"},
			{"ate", "eat", "tea"},
		}*/
		groupAnagrams(data)
	}
}
