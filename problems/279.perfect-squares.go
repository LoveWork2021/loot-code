package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/perfect-squares

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := numSquares(12)
	fmt.Println(r)
}

// COMMENT: 动态规划，多个数相加成一个数，最小组合
