package main

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	s := []int{}

	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	n := len(s) - 1
	s1 := make([]int, n+1)
	for i := 0; i <= n; i++ {
		s1[i] = s[n-i]
	}
	return s1
}
