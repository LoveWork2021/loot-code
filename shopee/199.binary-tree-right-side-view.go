package main

// https://leetcode-cn.com/problems/binary-tree-right-side-view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	vs := []int{}

	ns := []*TreeNode{root}
	for len(ns) > 0 {
		vs = append(vs, ns[0].Val)

		ns2 := []*TreeNode{}
		for _, n := range ns {
			if n.Right != nil {
				ns2 = append(ns2, n.Right)
			}

			if n.Left != nil {
				ns2 = append(ns2, n.Left)
			}
		}

		ns = ns2
	}

	return vs
}
