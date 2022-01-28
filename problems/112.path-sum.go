package main

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

}
