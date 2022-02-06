package main

import "fmt"

// https://leetcode-cn.com/problems/path-sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	vs := []int{}

	ns := []*TreeNode{root}
	for {
		_ns := []*TreeNode{}

		for _, n := range ns {
			hasLeft := n.Left != nil
			hasRight := n.Right != nil

			if hasLeft {
				n.Left.Val += n.Val
				_ns = append(_ns, n.Left)
			}

			if hasRight {
				n.Right.Val += n.Val
				_ns = append(_ns, n.Right)
			}

			if !hasLeft && !hasRight {
				vs = append(vs, n.Val)
			}
		}

		if len(_ns) == 0 {
			break
		}

		ns = _ns
	}

	for _, v := range vs {
		if v == targetSum {
			return true
		}
	}

	return false
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

	r := hasPathSum(toTree([]*V{&V{1}, &V{2}, &V{3}}), 5)
	fmt.Println(r)
}
