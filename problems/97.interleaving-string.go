package main

import "fmt"

// https://leetcode-cn.com/problems/interleaving-string

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)

	if l1+l2 != l3 {
		return false
	}

	f := make([][]bool, l1+1)
	for i := 0; i < l1; i++ {
		f[i] = make([]bool, l2+1)
	}

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			k := i + j - 1

			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[k])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s1[j-1] == s3[k])
			}
		}
	}

	return f[l1][l2]
}

func main() {
	r := isInterleave(
		"aabcc",
		"dbbca",
		"aadbbcbcac",
	)
	fmt.Println(r)
}
