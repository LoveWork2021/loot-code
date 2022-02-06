package main

import "fmt"

// https://leetcode-cn.com/problems/fibonacci-number

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 0, 1
	for i := 2; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

func main() {
	r := fib(2)
	fmt.Println(r)
}
