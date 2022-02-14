package main

// https://leetcode-cn.com/problems/add-two-numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}

	next := head
	adva := false
	for {
		if l1 == nil && l2 == nil && !adva {
			break
		}

		v := 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if adva {
			v += 1
		}

		if v > 9 {
			adva = true
			v = v - 10
		} else {
			adva = false
		}

		next.Next = &ListNode{Val: v}
		next = next.Next
	}

	return head.Next
}
