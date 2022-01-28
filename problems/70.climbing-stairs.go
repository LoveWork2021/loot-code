package main

import "fmt"

// https://leetcode-cn.com/problems/climbing-stairs

func climbStairs(n int) int {
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

func main() {
	r := climbStairs(2)
	fmt.Println(r)
}

// COMMENT: 动态规划！！！最后一步可能是两步或者一步，所以 f(x) = f(x-1) + f(x-2)
