package main

import "fmt"

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	ldr(root, &arr)
	return arr[k-1]
}

func ldr(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	ldr(node.Left, arr)
	*arr = append(*arr, node.Val)
	ldr(node.Right, arr)
}

func main() {
	type V struct {
		i int
	}

	toTree := func(list []*V) *TreeNode {
		ns := make([]*TreeNode, len(list))
		for i, v := range list {
			if v == nil {
				continue
			}

			j := i / 2

			n := &TreeNode{v.i, nil, nil}
			ns[i] = n

			if i != 0 && i%2 == 0 {
				j--
				ns[j].Right = n
			} else {
				ns[j].Left = n
			}
		}

		return ns[0]
	}

	r := kthSmallest(toTree([]*V{&V{3}, &V{1}, &V{34}, nil, &V{2}}), 1)
	fmt.Println(r)
}
