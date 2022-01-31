package main

import "fmt"

// https://leetcode-cn.com/problems/implement-strstr

// 暴力算法实现
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

func main() {
	r := strStr(
		"hello",
		"ll",
	)
	fmt.Println(r)
}
