package main

import "fmt"

// https://leetcode-cn.com/problems/valid-parentheses

func isValid(s string) bool {
	m := map[rune]int{
		'(': 1,
		')': -1,
		'{': 2,
		'}': -2,
		'[': 3,
		']': -3,
	}

	l := []int{}

	for i, c := range s {
		v, ok := m[c]
		if !ok {
			return false
		}

		if v > 0 {
			l = append(l, v)
		} else {
			n := len(l)
			if n == 0 {
				return false
			}

			v = -v

			if l[n-1] == v {
				l = l[:n-1]
			} else {
				return false
			}
		}
	}

	return len(l) == 0
}

func main() {
	r := isValid("()")
	fmt.Println(r)
}
