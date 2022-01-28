package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum

func threeSum(nums []int) [][]int {
	// 注意数组长度
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	l := len(nums)

	r := [][]int{}

	for ai := 0; ai < l; ai++ {
		a := nums[ai]

		// 因为数组排序，a > 0 说明后面不可能有大于 0 的组合
		if a > 0 {
			break
		}

		if ai > 0 && nums[ai-1] == a {
			continue
		}

		ci := l - 1

		for bi := ai + 1; bi < l; bi++ {
			b := nums[bi]

			if bi > ai+1 && nums[bi-1] == b {
				continue
			}

			for bi < ci && a+b+nums[ci] > 0 {
				ci--
			}

			if bi == ci {
				break
			}

			if a+b+nums[ci] == 0 {
				r = append(r, []int{a, b, nums[ci]})
			}
		}
	}

	return r
}

func main() {
	r := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(r)
}

// COMMENT: 先排序数组，可知 a <= b <= c，注意数组长度
