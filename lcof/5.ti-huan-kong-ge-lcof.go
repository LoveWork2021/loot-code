package main

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof

func replaceSpace(s string) string {
	s1 := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s1 += "%20"
		} else {
			s1 += string(s[i])
		}
	}

	return s1
}
