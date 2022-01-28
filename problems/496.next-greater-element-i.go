package main

import "fmt"

// https://leetcode-cn.com/problems/next-greater-element-i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}

	l := []int{}
	for _, n := range nums2 {
		for i := 0; i < len(l); {
			if l[i] < n {
				m[l[i]] = n
				// 删除中间元素
				l = append(l[:i], l[i+1:]...)
				continue
			}

			i++
		}

		l = append(l, n)
	}

	r := []int{}
	for _, n := range nums1 {
		if v, ok := m[n]; ok {
			r = append(r, v)
		} else {
			r = append(r, -1)
		}
	}
	return r
}

func main() {
	r := nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7})
	fmt.Println(r)
}
