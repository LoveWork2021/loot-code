package main

import "fmt"

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted

func twoSum(numbers []int, target int) []int {
	l, h := 0, len(numbers)-1
	for l < h {
		n := numbers[l] + numbers[h]

		if n == target {
			return []int{l + 1, h + 1}
		} else if n < target {
			l++
		} else {
			h--
		}
	}

	return []int{}
}

func main() {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(r)
}

// COMMENT: 双指针解法，利用数组有序
