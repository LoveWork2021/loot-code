package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/lfu-cache

type Node struct {
	key   int
	value int
	freq  int
	prev  *Node
	next  *Node
}

type Link struct {
	head *Node
	tail *Node
}

func NewLink() *Link {
	head := &Node{-1, -1, -1, nil, nil}
	tail := &Node{-2, -1, -1, nil, nil}

	head.next = tail
	tail.prev = head

	return &Link{head, tail}
}

func (this *Link) IsEmpty() bool {
	return this.head.next == this.tail
}

func (this *Link) AddHead(n *Node) {
	if n == this.head || n == this.tail {
		return
	}

	n.prev = this.head
	n.next = this.head.next

	this.head.next.prev = n
	this.head.next = n
}

func (this *Link) Remove(n *Node) {
	if n == this.head || n == this.tail || n.prev == nil {
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *Link) RemoveTail() *Node {
	if this.IsEmpty() {
		return nil
	}

	n := this.tail.prev
	this.Remove(n)
	return n
}

type LFUCache struct {
	capacity int
	keyMap   map[int]*Node
	freqMap  map[int]*Link
	minFreq  int
}

func (this LFUCache) getFreqLink(freq int) *Link {
	if link, ok := this.freqMap[freq]; ok {
		return link
	} else {
		link = NewLink()
		this.freqMap[freq] = link
		return link
	}
}

func (this *LFUCache) addFreq(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev

	if this.freqMap[n.freq].IsEmpty() {
		delete(this.freqMap, n.freq)

		if this.minFreq == n.freq {
			this.minFreq += 1
		}
	}

	link := this.getFreqLink(n.freq + 1)
	link.AddHead(n)

	n.freq += 1
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}

	var ok bool

	var n *Node
	if n, ok = this.keyMap[key]; !ok {
		return -1
	}

	this.addFreq(n)

	return n.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if n, ok := this.keyMap[key]; ok {
		n.value = value
		this.addFreq(n)
		return
	}

	if len(this.keyMap) >= this.capacity {
		link := this.freqMap[this.minFreq]

		n := link.RemoveTail()
		if n != nil {
			delete(this.keyMap, n.key)
		}

		if link.IsEmpty() {
			delete(this.freqMap, this.minFreq)
		}
	}

	n := &Node{key, value, 1, nil, nil}
	this.keyMap[key] = n

	link := this.getFreqLink(1)
	link.AddHead(n)

	this.minFreq = 1
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity,
		make(map[int]*Node, capacity),
		map[int]*Link{},
		0,
	}
}

func main() {
	arg := [][]int{
		{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4},
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
