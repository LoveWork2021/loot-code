package main

import "fmt"

// https://leetcode-cn.com/problems/perfect-squares

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
    // 最坏情况，全部为 1 相加 
		minn := i
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j]+1)
		}
		f[i] = minn
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
