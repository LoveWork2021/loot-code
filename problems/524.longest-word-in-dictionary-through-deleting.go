package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		a := dictionary[i]
		b := dictionary[j]

		if len(a) == len(b) {
			return a < b
		} else {
			return len(a) > len(b)
		}
	})

	for _, word := range dictionary {
		i := 0
		for _, r := range s {
			if word[i] == byte(r) {
				i++

				if i == len(word) {
					return word
				}

				continue
			}
		}
	}

	return ""
}

func main() {
	r := findLongestWord("aaa", []string{"aaa", "aa", "a"})
	fmt.Println(r)
}
