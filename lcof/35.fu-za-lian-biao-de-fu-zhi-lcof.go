package main

// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}

	var prev *Node
	next := head
	for next != nil {
		n := &Node{next.Val, nil, nil}
		m[next] = n

		if prev != nil {
			prev.Next = n
		}

		prev = n
		next = next.Next
	}

	next = head
	for next != nil {
		if next.Random != nil {
			m[next].Random = m[next.Random]
		}
		next = next.Next
	}

	return m[head]
}

func main() {

}
