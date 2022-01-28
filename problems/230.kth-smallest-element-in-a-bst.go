package main

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

}

// COMMENT: 二叉搜索树的左/右子树均小于/大于根节点；子树也是一颗二叉搜索树；中序遍历后是有序的数组
