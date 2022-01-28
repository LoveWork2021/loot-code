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

/*
 * COMMENT:
 * 栈：后进先出
 * 栈顶：插入删除数据的地方，如果使用数组，一般是数组的尾部为栈顶
 * 单调递减栈：（不同人不同定义）从栈底到栈顶的元素是单调递减，最大的元素在底部。假设栈 [2]，如果现在要插入 4，必须先删除 2，如果 4 入栈了，现在栈为 [2,4]，从底到顶是 2 -> 4，不符合递减
 *    每个弹出的元素，栈中都存在一个比它大的元素
 *
 * 把问题看成 acb
 * 单调递减栈保证了栈中存在元素大于当前出栈的元素
 * 同时从右向左（i--）遍历
 *     前两次直接入栈
 *     如果栈有弹出元素，说明存在 cb，因为先入栈的元素是 c，被弹出说明 b > c
 * 		 记录每次遍历最后一个弹出的元素，该元素会越来越大，如果该元素变小，说明早已存在 acb，比如 [4,5,6,3,4]
 *     如果一直没有弹出元素，表示不存在 cb 的情况
 *     如果有弹出元素，且此时 nums[i] < 弹出元素，那么符合 acb 的情况
 */
