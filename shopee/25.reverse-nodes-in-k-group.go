package main

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	prevHead := &ListNode{Next: head}

	// 当前组的上一个
	prevGroup := prevHead
	prev := prevHead

	node := head
	for i := 1; ; i++ {
		if i%k == 1 {
			prevGroup = prev
		}

		if i%k == 0 {
			// 交换 prevGroup -> a -> b -> c(node) -> next
			a := node.Next
			b := prevGroup.Next
			for j := 0; j < k; j++ {
				c := b.Next
				b.Next = a
				a = b
				b = c
			}
			node = prevGroup.Next
			prevGroup.Next = a
		}

		if node.Next == nil {
			return prevHead.Next
		} else {
			prev = node
			node = node.Next
		}
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5}

	head := &ListNode{Val: list[0]}
	node := head
	for i := 1; i < len(list); i++ {
		n := &ListNode{Val: list[i]}
		node.Next = n
		node = n
	}

	reverseKGroup(head, 2)
}
	node.Next = n
		node = n
	}

	reverseKGroup(head, 2)
}
