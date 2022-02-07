package main

// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof

type MinStack struct {
	stack1 []int
	stack2 []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	n := len(this.stack1)

	this.stack1 = append(this.stack1, x)

	if n == 0 || x < this.stack2[n-1] {
		this.stack2 = append(this.stack2, x)
	} else {
		this.stack2 = append(this.stack2, this.stack2[n-1])
	}
}

func (this *MinStack) Pop() {
	n := len(this.stack1)

	this.stack1 = this.stack1[:n-1]
	this.stack2 = this.stack2[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.stack1)

	if n == 0 {
		return -1
	} else {
		return this.stack1[0]
	}
}

func (this *MinStack) Min() int {
	n := len(this.stack1)

	if n == 0 {
		return -1
	} else {
		return this.stack2[n-1]
	}
}

func main() {

}
