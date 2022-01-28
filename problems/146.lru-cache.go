package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/lru-cache

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	keyMap   map[int]*Node
	head     *Node
	tail     *Node
}

func (this *LRUCache) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) addHeadNode(n *Node) {
	n.prev = this.head
	n.next = this.head.next

	this.head.next.prev = n
	this.head.next = n
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.keyMap[key]; !ok {
		return -1
	} else {
		this.removeNode(n)
		this.addHeadNode(n)
		return n.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.keyMap[key]; ok {
		n.value = value

		this.removeNode(n)
		this.addHeadNode(n)

		return
	}

	if len(this.keyMap) >= this.capacity {
		delete(this.keyMap, this.tail.prev.key)
		this.removeNode(this.tail.prev)
	}

	n := &Node{key, value, nil, nil}
	this.keyMap[key] = n
	this.addHeadNode(n)
}

func Constructor(capacity int) LRUCache {
	head := &Node{-1, -1, nil, nil}
	tail := &Node{-2, -1, nil, nil}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		keyMap:   make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func main() {
	arg := [][]int{
		{2},
		{1, 1},
		{2, 2},
		{1},
		{3, 3},
		{2},
		{4, 4},
		{1},
		{3},
		{4},
	}

	lru := Constructor(arg[0][0])

	for i := 1; i < len(arg); i++ {
		if len(arg[i]) == 1 {
			fmt.Printf("GET: %v:%v\n", arg[i][0], lru.Get(arg[i][0]))
		} else {
			lru.Put(arg[i][0], arg[i][1])
			fmt.Printf("PUT: %v:%v\n", arg[i][0], arg[i][1])
		}
	}
}
