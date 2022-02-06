package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/132-pattern

func find132pattern(nums []int) bool {
	l := len(nums)

	s := []int{nums[l-1]}

	out := math.MinInt64

	for i := l - 2; i >= 0; i-- {
		if nums[i] < out {
			return true
		}

		// 单调栈处理
		for len(s) > 0 && nums[i] > s[len(s)-1] {
			out = s[len(s)-1]
			s = s[:len(s)-1]
		}

		if nums[i] > out {
			s = append(s, nums[i])
		}
	}

	return false
}

func main() {
	r := find132pattern([]int{1, 3, 4, 5, 3})
	fmt.Println(r)
}
