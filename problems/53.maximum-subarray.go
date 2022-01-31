package main

import "fmt"

// https://leetcode-cn.com/problems/maximum-subarray

func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		// if nums[i]+nums[i-1] > nums[i] {
		// 	nums[i] = nums[i] + nums[i-1]
		// }

		if nums[i-1] > 0 {
			nums[i] = nums[i] + nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {
	r := maxSubArray([]int{1, 4, -1, 5, 7})
	fmt.Println(r)
}
