package main

// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{[]int{}, []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 && len(this.stack1) > 0 {
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = []int{}
	}

	n := len(this.stack2)
	if n == 0 {
		return -1
	}
	r := this.stack2[n-1]
	this.stack2 = this.stack2[:n-1]
	return r
}

func main() {

}
