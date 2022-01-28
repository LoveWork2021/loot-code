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

// COMMENT:
// 数组的某个连续子序列一定是以数组中某个值结尾，所以定义：
// f(i) 为以 arr[i] 结尾的连续子序列和的最大值
// 比如 [1,0,3]，f(2) 就是为以 3 结尾的连续子序列和的最大值
//
// 所以一个有 n 个元素数组的“最大连续子序列和”可以分解为 max{f(0), ...f(n-1)}
//
// 所以 f(i+1) = max{f(i) + arr[i+1], arr[i+1]}
