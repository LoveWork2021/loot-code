package main

import "fmt"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	max := 0
	set := map[rune]int{}

	for i, r := range s {
		if j, ok := set[r]; ok {
			if len(set) > max {
				max = len(set)
			}

			_set := map[rune]int{}
			for k, v := range set {
				if j < v {
					_set[k] = v
				}
			}
			set = _set
		}

		set[r] = i
	}

	if len(set) > max {
		return len(set)
	} else {
		return max
	}
}

func main() {
	r := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(r)
}
