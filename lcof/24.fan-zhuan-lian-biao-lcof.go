package main

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var a *ListNode = nil

	for head != nil {
		b := head.Next
		head.Next = a
		a = head
		head = b
	}

	return a
}
