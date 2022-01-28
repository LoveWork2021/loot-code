package main

import "fmt"

// https://leetcode-cn.com/problems/two-sum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		if n > target {
			continue
		}

		x := target - n
		xi, ok := m[x]
		if ok {
			return []int{i, xi}
		} else {
			m[n] = i
		}
	}

	return []int{}
}

func main() {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(r)
}
